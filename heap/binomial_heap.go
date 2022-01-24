package heap

import (
	"errors"
	"math"
)

type binomialTreeNode struct {
	element HeapElement
	order   int
	sibling *binomialTreeNode
	child   *binomialTreeNode
	parent  *binomialTreeNode
}

func newBinomialTree(el HeapElement) *binomialTreeNode {
	return &binomialTreeNode{
		element: el,
		order:   0,
		sibling: nil,
		child:   nil,
		parent:  nil,
	}
}

func (t *binomialTreeNode) link(t1 *binomialTreeNode) error {
	if t.order != t1.order {
		return errors.New("cannot link two binomial trees with different orders")
	}

	t1.parent = t
	t1.sibling = t.child
	t.child = t1
	t.order++
	return nil
}

type BinomialHeap struct {
	head     *binomialTreeNode
	heapType HeapType
	leftIsUp func(l, r HeapElement) bool
	cmp      func(l, r int) bool
	numEl    int
	nodeMap  map[int]*binomialTreeNode
}

func NewBinomialHeap(heapType HeapType) *BinomialHeap {
	h := &BinomialHeap{
		head:     nil,
		heapType: heapType,
		numEl:    0,
		nodeMap:  make(map[int]*binomialTreeNode),
	}

	if heapType == MIN_HEAP {
		h.leftIsUp = minHeapCmp
		h.cmp = func(l, r int) bool {
			return l < r
		}
	} else {
		h.leftIsUp = maxHeapCmp
		h.cmp = func(l, r int) bool {
			return l > r
		}
	}

	return h
}

func NewBinomialMinHeap() *BinomialHeap {
	return NewBinomialHeap(MIN_HEAP)
}

func NewBinomialMaxHeap() *BinomialHeap {
	return NewBinomialHeap(MAX_HEAP)
}

func merge(h1, h2 *BinomialHeap) *binomialTreeNode {
	if h2.head == nil {
		return h1.head
	}

	if h1.head == nil {
		return h2.head
	}

	// are these heaps mergeable?
	if h1.heapType != h2.heapType {
		return nil
	}

	// choose the root of the merged tree
	var r *binomialTreeNode
	n1 := h1.head
	n2 := h2.head
	if n1.order <= n2.order {
		r = n1
		n1 = r.sibling
	} else {
		r = n2
		n2 = r.sibling
	}

	aux := r
	for n1 != nil && n2 != nil {
		if n1.order <= n2.order {
			aux.sibling = n1
			n1 = n1.sibling
		} else {
			aux.sibling = n2
			n2 = n2.sibling
		}
		aux = aux.sibling
	}

	// what is left...
	if n1 != nil {
		aux.sibling = n1
	} else {
		aux.sibling = n2
	}

	return r
}

func (h *BinomialHeap) removeRoot(root, previous *binomialTreeNode) {
	if h.head == root {
		h.head = root.sibling
	} else {
		previous.sibling = root.sibling
	}

	var r *binomialTreeNode
	child := root.child
	for child != nil {
		next := child.sibling
		child.sibling = r
		child.parent = nil
		r = child
		child = next
	}

	tmpHeap := NewBinomialHeap(h.heapType)
	tmpHeap.head = r
	h.head = h.union(tmpHeap)
}

func (h *BinomialHeap) union(heap *BinomialHeap) *binomialTreeNode {
	r := merge(h, heap)
	h.head = nil
	heap.head = nil

	if r == nil {
		return nil
	}

	var prev *binomialTreeNode
	curr := r
	next := r.sibling
	for next != nil {
		if (curr.order != next.order) || (next.sibling != nil && next.sibling.order == curr.order) {
			prev = curr
			curr = next
		} else {
			if h.leftIsUp(curr.element, next.element) {
				curr.sibling = next.sibling
				curr.link(next)
			} else {
				if prev == nil {
					r = next
				} else {
					prev.sibling = next
				}
				next.link(curr)
				curr = next
			}
		}
		next = curr.sibling
	}

	return r
}

func (h *BinomialHeap) getTop(del bool) HeapElement {
	if h.head == nil {
		return nil
	}

	top := math.MaxInt64
	if h.heapType == MAX_HEAP {
		top = math.MinInt64
	}

	current := h.head
	var prevPtr, toRemove, toRemovePrev *binomialTreeNode
	var element HeapElement
	for current != nil {
		if h.cmp(current.element.Priority(), top) {
			top = current.element.Priority()
			element = current.element
			toRemove = current
			toRemovePrev = prevPtr
		}

		prevPtr = current
		current = current.sibling
	}

	if del {
		h.removeRoot(toRemove, toRemovePrev)
		h.numEl--
		delete(h.nodeMap, element.Priority())
	}
	return element
}

func (h *BinomialHeap) Peek() HeapElement {
	return h.getTop(false)
}

func (h *BinomialHeap) Pop() HeapElement {
	return h.getTop(true)
}

func (h *BinomialHeap) Push(el HeapElement) {
	t := newBinomialTree(el)
	hp := NewBinomialHeap(h.heapType)
	hp.head = t

	h.head = h.union(hp)
	h.numEl++
	h.nodeMap[el.Priority()] = t
}

func (h *BinomialHeap) Size() int {
	return h.numEl
}

func (h *BinomialHeap) Type() HeapType {
	return h.heapType
}

func (h *BinomialHeap) IsEmpty() bool {
	return h.numEl == 0
}

func (h *BinomialHeap) MoveUp(el HeapElement, newPriority int) error {
	// TODO check if the new priority is consistent with the heap type
	ptr, ok := h.nodeMap[el.Priority()]
	if !ok {
		return errors.New("no such element")
	}

	delete(h.nodeMap, el.Priority())
	ptr.element.SetPriority(newPriority)
	parent := ptr.parent
	for (parent != nil) && (h.leftIsUp(parent.element, ptr.element)) {
		// swap with its parent
		parentElement := parent.element
		parent.element = ptr.element
		ptr.element = parentElement

		h.nodeMap[parent.element.Priority()] = ptr

		ptr = parent
		parent = ptr.parent
	}
	h.nodeMap[newPriority] = ptr
	return nil
}

func (h *BinomialHeap) Merge(anotherHeap Heap) error {
	// TODO
	return errors.New("not implemented")
}
