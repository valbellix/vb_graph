package heap

import (
	"errors"
	"math"
)

type binaryHeap struct {
	array    []HeapElement
	heapType HeapType
}

func NewBinaryHeap(t HeapType) Heap {
	return &binaryHeap{
		array:    []HeapElement{},
		heapType: t,
	}
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
	if leftIndex < len(h.array) && leftIsUp(h, h.array[leftIndex], h.array[upmostIndex]) {
		upmostIndex = leftIndex
	}
	if rightIndex < len(h.array) && leftIsUp(h, h.array[rightIndex], h.array[upmostIndex]) {
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
	for leftIsUp(h, element, parentElement) {
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

func (h *binaryHeap) Remove(element HeapElement) {
	// TODO
}
