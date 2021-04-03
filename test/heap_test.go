package test

import (
	"testing"
	"vb_graph/heap"

	"github.com/stretchr/testify/assert"
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

func TestMinHeap(t *testing.T) {
	h := heap.NewBinaryHeap(heap.MIN_HEAP)

	h.Push(&testNode{50})
	h.Push(&testNode{1000})
	h.Push(&testNode{100})
	h.Push(&testNode{500})
	h.Push(&testNode{700})

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

	assert.Equal(t, len(expected), len(list))
	for i := range list {
		assert.Equal(t, expected[i].Priority(), list[i].Priority())
	}
}

func TestMaxHeap(t *testing.T) {
	h := heap.NewBinaryHeap(heap.MAX_HEAP)

	h.Push(&testNode{50})
	h.Push(&testNode{1000})
	h.Push(&testNode{100})
	h.Push(&testNode{500})
	h.Push(&testNode{700})

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

	assert.Equal(t, len(expected), len(list))
	for i := range list {
		assert.Equal(t, expected[i].Priority(), list[i].Priority())
	}
}
