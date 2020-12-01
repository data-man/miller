package transformers

import (
	"flag"
	"fmt"
	"os"

	"miller/clitypes"
	"miller/transforming"
	"miller/types"
)

// ----------------------------------------------------------------
var SeqgenSetup = transforming.TransformerSetup{
	Verb:         "seqgen",
	ParseCLIFunc: transformerSeqgenParseCLI,
	IgnoresInput: true,
}

func transformerSeqgenParseCLI(
	pargi *int,
	argc int,
	args []string,
	errorHandling flag.ErrorHandling, // ContinueOnError or ExitOnError
	_ *clitypes.TReaderOptions,
	__ *clitypes.TWriterOptions,
) transforming.IRecordTransformer {

	// Get the verb name from the current spot in the mlr command line
	argi := *pargi
	verb := args[argi]
	argi++

	// Parse local flags
	flagSet := flag.NewFlagSet(verb, errorHandling)

	pFieldName := flagSet.String(
		"f",
		"i",
		"Field name for counters.",
	)

	pStartString := flagSet.String(
		"start",
		"1",
		"Inclusive start value.",
	)

	pStopString := flagSet.String(
		"stop",
		"100",
		"Inclusive stop value.",
	)

	pStepString := flagSet.String(
		"step",
		"1",
		"Step value",
	)

	flagSet.Usage = func() {
		ostream := os.Stderr
		if errorHandling == flag.ContinueOnError { // help intentionally requested
			ostream = os.Stdout
		}
		transformerSeqgenUsage(ostream, args[0], verb, flagSet)
	}
	flagSet.Parse(args[argi:])
	if errorHandling == flag.ContinueOnError { // help intentionally requested
		return nil
	}

	// Find out how many flags were consumed by this verb and advance for the
	// next verb
	argi = len(args) - len(flagSet.Args())

	transformer, _ := NewTransformerSeqgen(
		*pFieldName,
		*pStartString,
		*pStopString,
		*pStepString,
	)

	*pargi = argi
	return transformer
}

func transformerSeqgenUsage(
	o *os.File,
	argv0 string,
	verb string,
	flagSet *flag.FlagSet,
) {
	fmt.Fprintf(o, "Usage: %s %s [options]\n", argv0, verb)
	fmt.Fprintf(o, "Passes input records directly to output. Most useful for format conversion.\n")
	fmt.Fprintf(o, "Produces a sequence of counters.  Discards the input record stream. Produces\n")
	fmt.Fprintf(o, "output as specified by the following options:\n")
	// flagSet.PrintDefaults() doesn't let us control stdout vs stderr
	flagSet.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(o, " -%v (default %v) %v\n", f.Name, f.Value, f.Usage) // f.Name, f.Value
	})
	fmt.Fprintf(o, "Start, stop, and/or step may be floating-point. Output is integer if start,\n")
	fmt.Fprintf(o, "stop, and step are all integers. Step may be negative. It may not be zero\n")
	fmt.Fprintf(o, "unless start == stop.\n")
}

// ----------------------------------------------------------------
type TransformerSeqgen struct {
	fieldName string
	start     types.Mlrval
	stop      types.Mlrval
	step      types.Mlrval
}

// ----------------------------------------------------------------
func NewTransformerSeqgen(
	fieldName string,
	startString string,
	stopString string,
	stepString string,
) (*TransformerSeqgen, error) {
	// xxx check int/float

	return &TransformerSeqgen{
		fieldName: fieldName,
		start:     types.MlrvalFromInferredType(startString),
		stop:      types.MlrvalFromInferredType(stopString),
		step:      types.MlrvalFromInferredType(stepString),
	}, nil
}

// ----------------------------------------------------------------
func (this *TransformerSeqgen) Map(
	inrecAndContext *types.RecordAndContext,
	outputChannel chan<- *types.RecordAndContext,
) {
	counter := this.start
	context := types.NewContext()
	context.UpdateForStartOfFile("seqgen")

	for {
		mdone := types.MlrvalGreaterThan(&counter, &this.stop)
		done, _ := mdone.GetBoolValue()
		if done {
			break
		}

		outrec := types.NewMlrmapAsRecord()
		outrec.PutCopy(&this.fieldName, &counter)

		context.UpdateForInputRecord(outrec)

		outrecAndContext := types.NewRecordAndContext(outrec, context)
		outputChannel <- outrecAndContext

		counter = types.MlrvalBinaryPlus(&counter, &this.step)
	}

	outrecAndContext := types.NewRecordAndContext(nil, context)
	outputChannel <- outrecAndContext
}