package heap

type HeapType int

const (
	MAX_HEAP HeapType = iota
	MIN_HEAP
)

type HeapElement interface {
	Priority() int
	SetPriority(p int)
}

type Heap interface {
	Size() int
	Push(el HeapElement)
	Pop() HeapElement
	IsEmpty() bool
	Type() HeapType
	MoveUp(el HeapElement, newPriority int) error
}

func leftIsUp(h Heap, elLeft, elRight HeapElement) bool {
	if h.Type() == MIN_HEAP {
		return elLeft.Priority() < elRight.Priority()
	} else {
		return elLeft.Priority() > elRight.Priority()
	}
}
