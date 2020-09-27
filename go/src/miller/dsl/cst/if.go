package cst

import (
	"errors"

	"miller/dsl"
	"miller/lib"
	"miller/types"
)

// ================================================================
// This is for if/elif/elif/else chains.
// ================================================================

// ----------------------------------------------------------------
type IfChainNode struct {
	ifItems []*IfItem
}

func NewIfChainNode(ifItems []*IfItem) *IfChainNode {
	return &IfChainNode{
		ifItems: ifItems,
	}
}

// ----------------------------------------------------------------
// For each if/elif/elif/else portion: the conditional part (...) and the
// statement-block part {...}. For "else", the conditional is nil.
type IfItem struct {
	conditionNode      IEvaluable
	statementBlockNode *StatementBlockNode
}

// ----------------------------------------------------------------
// Sample AST:

// DSL EXPRESSION:
// if (NR == 1) { $z = 100 } elif (NR == 2) { $z = 200 } elif (NR == 3) { $z = 300 } else { $z = 900 }
// AST:
// * StatementBlock
//     * IfChain
//         * IfItem "if"
//             * Operator "=="
//                 * ContextVariable "NR"
//                 * IntLiteral "1"
//             * StatementBlock
//                 * Assignment "="
//                     * DirectFieldValue "z"
//                     * IntLiteral "100"
//         * IfItem "elif"
//             * Operator "=="
//                 * ContextVariable "NR"
//                 * IntLiteral "2"
//             * StatementBlock
//                 * Assignment "="
//                     * DirectFieldValue "z"
//                     * IntLiteral "200"
//         * IfItem "elif"
//             * Operator "=="
//                 * ContextVariable "NR"
//                 * IntLiteral "3"
//             * StatementBlock
//                 * Assignment "="
//                     * DirectFieldValue "z"
//                     * IntLiteral "300"
//         * IfItem "else"
//             * StatementBlock
//                 * Assignment "="
//                     * DirectFieldValue "z"
//                     * IntLiteral "900"

func BuildIfChainNode(astNode *dsl.ASTNode) (*IfChainNode, error) {
	lib.InternalCodingErrorIf(astNode.Type != dsl.NodeTypeIfChain)

	ifItems := make([]*IfItem, 0)

	astChildren := astNode.Children

	for _, astChild := range astChildren {
		lib.InternalCodingErrorIf(astChild.Type != dsl.NodeTypeIfItem)
		token := string(astChild.Token.Lit) // "if", "elif", "else"
		if token == "if" || token == "elif" {
			lib.InternalCodingErrorIf(len(astChild.Children) != 2)
			conditionNode, err := BuildEvaluableNode(astChild.Children[0])
			if err != nil {
				return nil, err
			}
			statementBlockNode, err := BuildStatementBlockNode(astChild.Children[1])
			if err != nil {
				return nil, err
			}
			ifItem := &IfItem{
				conditionNode:      conditionNode,
				statementBlockNode: statementBlockNode,
			}
			ifItems = append(ifItems, ifItem)

		} else if token == "else" {
			lib.InternalCodingErrorIf(len(astChild.Children) != 1)
			var conditionNode IEvaluable = nil
			statementBlockNode, err := BuildStatementBlockNode(astChild.Children[0])
			if err != nil {
				return nil, err
			}
			ifItem := &IfItem{
				conditionNode:      conditionNode,
				statementBlockNode: statementBlockNode,
			}
			ifItems = append(ifItems, ifItem)

		} else {
			lib.InternalCodingErrorIf(true)
		}
	}
	ifChainNode := NewIfChainNode(ifItems)
	return ifChainNode, nil
}

// ----------------------------------------------------------------
func (this *IfChainNode) Execute(state *State) error {
	for _, ifItem := range this.ifItems {
		condition := types.MlrvalFromTrue()
		if ifItem.conditionNode != nil {
			condition = ifItem.conditionNode.Evaluate(state)
		}
		boolValue, isBool := condition.GetBoolValue()
		if !isBool {
			// TODO: line-number/token info for the DSL expression.
			return errors.New("Miller: conditional expression did not evaluate to boolean.")
		}
		if boolValue == true {
			err := ifItem.statementBlockNode.Execute(state)
			if err != nil {
				return err
			}
			break
		}
	}
	// TODO
	return nil
}