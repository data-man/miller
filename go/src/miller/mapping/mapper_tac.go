package mapping

import (
	"container/list"

	"miller/containers"
	"miller/runtime"
)

type MapperTac struct {
	lrecs *list.List
}

func NewMapperTac() (*MapperTac, error) {
	return &MapperTac{
		lrecs: list.New(),
	}, nil
}

func (this *MapperTac) Map(
	inrec *containers.Lrec,
	context *runtime.Context,
	outrecs chan<- *containers.Lrec,
) {
	if inrec != nil {
		this.lrecs.PushFront(inrec)
	} else {
		// end of stream
		for e := this.lrecs.Front(); e != nil; e = e.Next() {
			outrecs <- e.Value.(*containers.Lrec)
		}
		outrecs <- nil
	}
}
