package input

import (
	"miller/containers"
)

// Since Go is concurrent, the context struct (AWK-like variables such as
// FILENAME, NF, NF, FNR, etc.) needs to be duplicated and passed through the
// channels along with each record. Hence the initial context, which readers
// update on each new file/record, and the channel of containers.LrecAndContext
// rather than channel of containers.Lrec.

type IRecordReader interface {
	Read(
		filenames []string,
		initialContext containers.Context,
		inrecsAndContexts chan<- *containers.LrecAndContext,
		echan chan error,
	)
}
