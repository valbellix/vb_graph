package heap

import (
	"errors"
	"math"
)

type binaryHeap struct {
	array    []HeapElement
	heapType HeapType
	leftIsUp func(HeapElement, HeapElement) bool
}

func NewBinaryHeap(t HeapType) Heap {
	h := &binaryHeap{
		array:    []HeapElement{},
		heapType: t,
	}

	if t == MIN_HEAP {
		h.leftIsUp = func(l, r HeapElement) bool {
			return l.Priority() < r.Priority()
		}
	} else {
		h.leftIsUp = func(l, r HeapElement) bool {
			return l.Priority() > r.Priority()
		}
	}

	return h
}

func NewBinaryMinHeap() Heap {
	return NewBinaryHeap(MIN_HEAP)
}

func NewBinaryMaxHeap() Heap {
	return NewBinaryHeap(MAX_HEAP)
}

//////
// These functions will be helpful to manage trees implemented as a vector-like structure

func (h *binaryHeap) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *binaryHeap) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h *binaryHeap) parentIndex(index int) (int, error) {
	if index == 0 {
		return 0, errors.New("error: the root has no parent")
	}
	return int(math.Floor((float64(index) - 1) / 2)), nil
}

func (h *binaryHeap) swapElements(i, j int) {
	aux := h.array[i]
	h.array[i] = h.array[j]
	h.array[j] = aux
}

func (h *binaryHeap) heapify(i int) {
	leftIndex := h.leftChildIndex(i)
	rightIndex := h.rightChildIndex(i)

	upmostIndex := i
	if leftIndex < len(h.array) && h.leftIsUp(h.array[leftIndex], h.array[upmostIndex]) {
		upmostIndex = leftIndex
	}
	if rightIndex < len(h.array) && h.leftIsUp(h.array[rightIndex], h.array[upmostIndex]) {
		upmostIndex = rightIndex
	}
	if upmostIndex != i {
		h.swapElements(i, upmostIndex)
		h.heapify(upmostIndex)
	}
}

//////

func (h *binaryHeap) Size() int {
	return len(h.array)
}

func (h *binaryHeap) Type() HeapType {
	return h.heapType
}

func (h *binaryHeap) IsEmpty() bool {
	return len(h.array) == 0
}

func (h *binaryHeap) Push(element HeapElement) {
	// this will be the index of the element to append in the list
	current := len(h.array)

	h.array = append(h.array, element)
	parent, err := h.parentIndex(current)

	// if the current node is the root, we just return
	if err != nil {
		return
	}

	parentElement := h.array[parent]
	for h.leftIsUp(element, parentElement) {
		h.swapElements(current, parent)

		current = parent
		parent, err = h.parentIndex(current)
		if err != nil {
			break
		}
		parentElement = h.array[parent]
	}
}

func (h *binaryHeap) Pop() HeapElement {
	if h.IsEmpty() {
		return nil
	}

	top := h.array[0]
	size := len(h.array)
	if size == 1 {
		// we are retaining the capacity to avoid reallocation later if we need to reuse
		h.array = h.array[:0]
		return top
	}

	// swap the root with the latest element of the array and heapify
	h.array[0] = h.array[size-1]
	h.array = h.array[:size-1]

	h.heapify(0)
	return top
}

// it now scans the array... we need a map to efficiently handle this
func (h *binaryHeap) getIndex(element HeapElement) (int, error) {
	for i, v := range h.array {
		if v.Priority() == element.Priority() {
			return i, nil
		}
	}

	return -1, errors.New("error: element is not in the heap")
}

func (h *binaryHeap) MoveUp(element HeapElement, newPriority int) error {
	index, err := h.getIndex(element)
	if err != nil {
		return err
	}

	h.array[index].SetPriority(newPriority)
	parent := int(math.Floor((float64(index - 1)) / 2))
	for index > 0 && h.leftIsUp(h.array[index], h.array[parent]) {
		h.swapElements(index, parent)
		index = parent
		parent = int(math.Floor((float64(index - 1)) / 2))
	}

	return nil
}

func (h *binaryHeap) Merge(anotherHeap Heap) error {
	if h.Type() != anotherHeap.Type() {
		return errors.New("the heap to merge is of a different type")
	}
	for !anotherHeap.IsEmpty() {
		h.Push(anotherHeap.Pop())
	}

	return nil
}
