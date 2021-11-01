package test

import (
	"testing"
	"vb_graph/heap"

	"github.com/stretchr/testify/assert"
)

type HeapFamily int

const (
	BINARY_HEAP HeapFamily = iota
	BINOMIAL_HEAP
)

type testNode struct {
	priority int
}

func (t *testNode) Priority() int {
	return t.priority
}

func (t *testNode) SetPriority(p int) {
	t.priority = p
}

func makeAListOfHeapElements(h heap.Heap) []heap.HeapElement {
	l := []heap.HeapElement{}

	for !h.IsEmpty() {
		l = append(l, h.Pop())
	}

	return l
}

func checkListContent(t *testing.T, expected, list []heap.HeapElement) {
	assert.Equal(t, len(expected), len(list))
	for i := range list {
		assert.Equal(t, expected[i].Priority(), list[i].Priority())
	}
}

func buildTestHeap(f HeapFamily, t heap.HeapType) heap.Heap {
	var h heap.Heap
	if f == BINARY_HEAP {
		h = heap.NewBinaryHeap(t)
	} else if f == BINOMIAL_HEAP {
		h = heap.NewBinomialHeap(t)
	}

	h.Push(&testNode{50})
	h.Push(&testNode{1000})
	h.Push(&testNode{100})
	h.Push(&testNode{500})
	h.Push(&testNode{700})

	return h
}

func buildEmptyHeap(f HeapFamily, t heap.HeapType) heap.Heap {
	var h heap.Heap
	if f == BINARY_HEAP {
		h = heap.NewBinaryHeap(t)
	} else if f == BINOMIAL_HEAP {
		h = heap.NewBinomialHeap(t)
	}

	return h
}

func TestMinBinaryHeap(t *testing.T) {
	testMinHeap(t, BINARY_HEAP)
}

func TestMinBinomialHeap(t *testing.T) {
	testMinHeap(t, BINOMIAL_HEAP)
}

func TestMaxBinaryHeap(t *testing.T) {
	testMaxHeap(t, BINARY_HEAP)
}

func TestMaxBinomialHeap(t *testing.T) {
	testMaxHeap(t, BINOMIAL_HEAP)
}

func testMinHeap(t *testing.T, family HeapFamily) {
	h := buildTestHeap(family, heap.MIN_HEAP)

	assert.Equal(t, 5, h.Size())

	n := h.Pop()
	assert.Equal(t, 4, h.Size())
	assert.Equal(t, 50, n.Priority())

	list := make([]heap.HeapElement, 0)
	list = append(list, n)

	for !h.IsEmpty() {
		list = append(list, h.Pop())
	}
	assert.True(t, h.IsEmpty())

	expected := []heap.HeapElement{
		&testNode{50}, &testNode{100}, &testNode{500}, &testNode{700}, &testNode{1000},
	}
	checkListContent(t, expected, list)

	h = buildTestHeap(family, heap.MIN_HEAP)

	err := h.MoveUp(&testNode{700}, 10)
	assert.Nil(t, err, "error should not occur")

	list = makeAListOfHeapElements(h)

	expected = []heap.HeapElement{
		&testNode{10}, &testNode{50}, &testNode{100}, &testNode{500}, &testNode{1000},
	}

	checkListContent(t, expected, list)

	h1 := buildEmptyHeap(family, heap.MIN_HEAP)
	h1.Push(&testNode{10})
	h1.Push(&testNode{500})
	h1.Push(&testNode{1000})

	h2 := buildEmptyHeap(family, heap.MIN_HEAP)
	h2.Push(&testNode{50})
	h2.Push(&testNode{100})

	err = h1.Merge(h2)
	assert.Nil(t, err, "error should not occur")
	assert.Equal(t, 5, h1.Size(), "size should be sum of h1+h2")
	assert.True(t, h2.IsEmpty(), true, "h2 should be empty")

	list = makeAListOfHeapElements(h1)
	checkListContent(t, expected, list)
}

func testMaxHeap(t *testing.T, family HeapFamily) {
	h := buildTestHeap(family, heap.MAX_HEAP)

	assert.Equal(t, 5, h.Size())

	n := h.Pop()
	assert.Equal(t, 4, h.Size())
	assert.Equal(t, 1000, n.Priority())

	list := make([]heap.HeapElement, 0)
	list = append(list, n)

	for !h.IsEmpty() {
		list = append(list, h.Pop())
	}
	assert.True(t, h.IsEmpty())

	expected := []heap.HeapElement{
		&testNode{1000}, &testNode{700}, &testNode{500}, &testNode{100}, &testNode{50},
	}
	checkListContent(t, expected, list)

	h = buildTestHeap(family, heap.MAX_HEAP)

	err := h.MoveUp(&testNode{100}, 1500)
	assert.Nil(t, err, "error should not occur")

	list = makeAListOfHeapElements(h)

	expected = []heap.HeapElement{
		&testNode{1500}, &testNode{1000}, &testNode{700}, &testNode{500}, &testNode{50},
	}
	checkListContent(t, expected, list)

	h1 := buildEmptyHeap(family, heap.MAX_HEAP)
	h1.Push(&testNode{1500})
	h1.Push(&testNode{500})
	h1.Push(&testNode{50})

	h2 := buildEmptyHeap(family, heap.MAX_HEAP)
	h1.Push(&testNode{1000})
	h1.Push(&testNode{700})

	err = h1.Merge(h2)
	assert.Nil(t, err, "error should not occur")
	assert.Equal(t, 5, h1.Size(), "size should be sum of h1+h2")
	assert.True(t, h2.IsEmpty(), true, "h2 should be empty")

	list = makeAListOfHeapElements(h1)
	checkListContent(t, expected, list)
}
