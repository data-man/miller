// ================================================================
// This handles print, printn, eprint, and eprintn statements.
// ================================================================

package cst

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"miller/dsl"
	"miller/lib"
	"miller/output"
	"miller/types"
)

// ----------------------------------------------------------------
// Example ASTs:
//
// $ mlr -n put -v 'print $a, $b'
// DSL EXPRESSION:
// print $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * no-op
//
// $ mlr -n put -v 'print > stdout, $a, $b'
// DSL EXPRESSION:
// print > stdout, $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * redirect write ">"
//             * stdout redirect target "stdout"
//
// $ mlr -n put -v 'print > stderr, $a, $b'
// DSL EXPRESSION:
// print > stderr, $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * redirect write ">"
//             * stderr redirect target "stderr"
//
// $ mlr -n put -v 'print > "foo.dat", $a, $b'
// DSL EXPRESSION:
// print > "foo.dat", $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * redirect write ">"
//             * string literal "foo.dat"
//
// $ mlr -n put -v 'print >> "foo.dat", $a, $b'
// DSL EXPRESSION:
// print >> "foo.dat", $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * redirect append ">>"
//             * string literal "foo.dat"
//
// $ mlr -n put -v 'print | "command", $a, $b'
// DSL EXPRESSION:
// print | "command", $a, $b
// RAW AST:
// * statement block
//     * print statement "print"
//         * function callsite
//             * direct field value "a"
//             * direct field value "b"
//         * redirect pipe "|"
//             * string literal "command"
//
//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Corresponding data structures for these cases:
//
// * printToRedirectFunc is either printToStdout, printToStderr, or
//   printToFileOrPipe. Only the third of these takes a non-nil
//   redirectorTargetEvaluable and a non-nil outputHandlerManager.
//
// * redirectorTargetEvaluable is nil for stdout or stderr.
//
// * The OutputHandlerManager is for file names or commands in >, >> or |.
//   This is because the target for the redirect can vary from one record to
//   the next, e.g. mlr put 'print > $a.txt, $b'. The OutputHandlerManager
//   keeps file-handles for each distinct value of $a.
//
// So:
//
// * print $a, $b
//   AST redirectorNode         = NodeTypeNoOp
//     AST redirectorTargetNode = (none)
//   printToRedirectFunc        = printToStdout
//   redirectorTargetEvaluable  = nil
//   outputHandlerManager       = nil
//
// * print > stdout, $a, $b
//   AST redirectorNode         = NodeTypeRedirectWrite
//     AST redirectorTargetNode = NodeTypeRedirectTargetStdout
//   printToRedirectFunc        = printToStdout
//   redirectorTargetEvaluable  = nil
//   outputHandlerManager       = nil
//
// * print > stderr, $a, $b
//   AST redirectorNode         = NodeTypeRedirectWrite
//     AST redirectorTargetNode = NodeTypeRedirectTargetStderr
//   printToRedirectFunc        = printToStderr
//   redirectorTargetEvaluable  = nil
//   outputHandlerManager       = nil
//
// * print > "foo.dat", $a, $b
//   AST redirectorNode         = NodeTypeRedirectWrite
//     AST redirectorTargetNode = any of various evaluables
//   printToRedirectFunc        = printToFileOrPipe
//   redirectorTargetEvaluable  = non-nil
//   outputHandlerManager       = non-nil
//
// * print >> "foo.dat", $a, $b
//   AST redirectorNode         = NodeTypeRedirectAppend
//     AST redirectorTargetNode = any of various evaluables
//   printToRedirectFunc        = printToFileOrPipe
//   redirectorTargetEvaluable  = non-nil
//   outputHandlerManager       = non-nil
//
// * print | "command", $a, $b
//   AST redirectorNode         = NodeTypeRedirectPipe
//     AST redirectorTargetNode = any of various evaluables
//   printToRedirectFunc        = printToFileOrPipe
//   redirectorTargetEvaluable  = non-nil
//   outputHandlerManager       = non-nil

// ================================================================
type tPrintToRedirectFunc func(
	outputString string,
	state *State,
) error

type PrintStatementNode struct {
	expressionEvaluables      []IEvaluable
	terminator                string
	printToRedirectFunc       tPrintToRedirectFunc
	redirectorTargetEvaluable IEvaluable                  // for file/pipe targets
	outputHandlerManager      output.OutputHandlerManager // for file/pipe targets
}

// ----------------------------------------------------------------
func (this *RootNode) BuildPrintStatementNode(astNode *dsl.ASTNode) (IExecutable, error) {
	lib.InternalCodingErrorIf(astNode.Type != dsl.NodeTypePrintStatement)
	return this.buildPrintxStatementNode(
		astNode,
		os.Stdout,
		"\n",
	)
}

func (this *RootNode) BuildPrintnStatementNode(astNode *dsl.ASTNode) (IExecutable, error) {
	lib.InternalCodingErrorIf(astNode.Type != dsl.NodeTypePrintnStatement)
	return this.buildPrintxStatementNode(
		astNode,
		os.Stdout,
		"",
	)
}

func (this *RootNode) BuildEprintStatementNode(astNode *dsl.ASTNode) (IExecutable, error) {
	lib.InternalCodingErrorIf(astNode.Type != dsl.NodeTypeEprintStatement)
	return this.buildPrintxStatementNode(
		astNode,
		os.Stderr,
		"\n",
	)
}

func (this *RootNode) BuildEprintnStatementNode(astNode *dsl.ASTNode) (IExecutable, error) {
	lib.InternalCodingErrorIf(astNode.Type != dsl.NodeTypeEprintnStatement)
	return this.buildPrintxStatementNode(
		astNode,
		os.Stderr,
		"",
	)
}

// ----------------------------------------------------------------
// Common code for building print/eprint/printn/eprintn nodes

func (this *RootNode) buildPrintxStatementNode(
	astNode *dsl.ASTNode,
	defaultOutputStream *os.File,
	terminator string,
) (IExecutable, error) {
	lib.InternalCodingErrorIf(len(astNode.Children) != 2)
	expressionsNode := astNode.Children[0]
	redirectorNode := astNode.Children[1]

	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Things to be printed, e.g. $a and $b in 'print > "foo.dat", $a, $b'.

	var expressionEvaluables []IEvaluable = nil

	if expressionsNode.Type == dsl.NodeTypeNoOp {
		// Just 'print' without 'print $something'
		expressionEvaluables = make([]IEvaluable, 1)
		expressionEvaluable := this.BuildStringLiteralNode("")
		expressionEvaluables[0] = expressionEvaluable
	} else if expressionsNode.Type == dsl.NodeTypeFunctionCallsite {
		expressionEvaluables = make([]IEvaluable, len(expressionsNode.Children))
		for i, childNode := range expressionsNode.Children {
			expressionEvaluable, err := this.BuildEvaluableNode(childNode)
			if err != nil {
				return nil, err
			}
			expressionEvaluables[i] = expressionEvaluable
		}
	} else {
		lib.InternalCodingErrorIf(true)
	}

	//  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Redirection targets (the thing after > >> |, if any).

	retval := &PrintStatementNode{
		expressionEvaluables:      expressionEvaluables,
		terminator:                terminator,
		printToRedirectFunc:       nil,
		redirectorTargetEvaluable: nil,
		outputHandlerManager:      nil,
	}

	if redirectorNode.Type == dsl.NodeTypeNoOp {
		// No > >> or | was provided.
		if defaultOutputStream == os.Stdout {
			retval.printToRedirectFunc = retval.printToStdout
		} else if defaultOutputStream == os.Stderr {
			retval.printToRedirectFunc = retval.printToStderr
		} else {
			lib.InternalCodingErrorIf(true)
		}
	} else {
		// There is > >> or | provided.
		lib.InternalCodingErrorIf(redirectorNode.Children == nil)
		lib.InternalCodingErrorIf(len(redirectorNode.Children) != 1)
		redirectorTargetNode := redirectorNode.Children[0]
		var err error = nil

		if redirectorTargetNode.Type == dsl.NodeTypeRedirectTargetStdout {
			retval.printToRedirectFunc = retval.printToStdout
		} else if redirectorTargetNode.Type == dsl.NodeTypeRedirectTargetStderr {
			retval.printToRedirectFunc = retval.printToStderr
		} else {
			retval.printToRedirectFunc = retval.printToFileOrPipe

			retval.redirectorTargetEvaluable, err = this.BuildEvaluableNode(redirectorTargetNode)
			if err != nil {
				return nil, err
			}

			if redirectorNode.Type == dsl.NodeTypeRedirectWrite {
				retval.outputHandlerManager = output.NewFileWritetHandlerManager(this.recordWriterOptions)
			} else if redirectorNode.Type == dsl.NodeTypeRedirectAppend {
				retval.outputHandlerManager = output.NewFileAppendHandlerManager(this.recordWriterOptions)
			} else if redirectorNode.Type == dsl.NodeTypeRedirectPipe {
				retval.outputHandlerManager = output.NewPipeWriteHandlerManager(this.recordWriterOptions)
			} else {
				return nil, errors.New(
					fmt.Sprintf(
						"%s: unhandled redirector node type %s.",
						os.Args[0], string(redirectorNode.Type),
					),
				)
			}
		}
	}

	// Register this with the CST root node so that open file descriptrs can be
	// closed, etc at end of stream.
	if retval.outputHandlerManager != nil {
		this.RegisterOutputHandlerManager(retval.outputHandlerManager)
	}

	return retval, nil
}

// ----------------------------------------------------------------
func (this *PrintStatementNode) Execute(state *State) (*BlockExitPayload, error) {
	if len(this.expressionEvaluables) == 0 {
		this.printToRedirectFunc(this.terminator, state)
	} else {
		// 5x faster than fmt.Print() separately: note that os.Stdout is
		// non-buffered in Go whereas stdout is buffered in C.
		//
		// Minus: we need to do our own buffering for performance.
		//
		// Plus: we never have to worry about forgetting to do fflush(). :)
		var buffer bytes.Buffer

		for i, expressionEvaluable := range this.expressionEvaluables {
			if i > 0 {
				buffer.WriteString(" ")
			}
			evaluation := expressionEvaluable.Evaluate(state)
			if !evaluation.IsAbsent() {
				buffer.WriteString(evaluation.String())
			}
		}
		buffer.WriteString(this.terminator)
		this.printToRedirectFunc(buffer.String(), state)
	}
	return nil, nil
}

// ----------------------------------------------------------------
func (this *PrintStatementNode) printToStdout(
	outputString string,
	state *State,
) error {
	// Insert the string into the record-output stream, so that goroutine can
	// print it, resulting in deterministic output-ordering.
	state.OutputChannel <- types.NewOutputString(outputString, state.Context)
	return nil
}

// ----------------------------------------------------------------
func (this *PrintStatementNode) printToStderr(
	outputString string,
	state *State,
) error {
	fmt.Fprint(os.Stderr, outputString)
	return nil
}

// ----------------------------------------------------------------
func (this *PrintStatementNode) printToFileOrPipe(
	outputString string,
	state *State,
) error {
	redirectorTarget := this.redirectorTargetEvaluable.Evaluate(state)
	if !redirectorTarget.IsString() {
		return errors.New(
			fmt.Sprintf(
				"%s: output redirection yielded %s, not string.",
				os.Args[0], redirectorTarget.GetTypeName(),
			),
		)
	}
	outputFileName := redirectorTarget.String()

	this.outputHandlerManager.WriteString(outputString, outputFileName)
	return nil
}
