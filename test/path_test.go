package test

import (
	"testing"
	"vb_graph/graph"
	"vb_graph/heap"

	"github.com/stretchr/testify/assert"
)

func buildDistToVerify(distances map[graph.Node]int) map[string]int {
	m := make(map[string]int)
	for k, v := range distances {
		m[k.GetLabel()] = v
	}
	return m
}

func buildPrevToVerify(previous map[graph.Node]graph.Node) map[string]string {
	m := make(map[string]string)
	for k, v := range previous {
		l := ""
		if v != nil {
			l = v.GetLabel()
		}
		m[k.GetLabel()] = l
	}
	return m
}

func buildExpectedDistances() map[string]int {
	m := make(map[string]int)
	m["A"] = 0
	m["B"] = 4
	m["C"] = 4
	m["D"] = 7
	m["F"] = 8
	m["E"] = 5
	return m
}

func buildExpectedPrevious() map[string]string {
	m := make(map[string]string)
	m["E"] = "C"
	m["A"] = ""
	m["B"] = "A"
	m["C"] = "A"
	m["D"] = "C"
	m["F"] = "E"
	return m
}

func testShortestPath(t *testing.T, h heap.Heap) {
	g, err := graph.ParseDIMACS("data/test_path.dimacs", "e")
	assert.Nil(t, err, "ParseDIMACS should not return any error")

	distances, previous, err := graph.ShortestPath(g, g.GetRoot(), h)
	assert.Nil(t, err)
	assert.NotZero(t, len(distances), "distances should not be empty")
	assert.NotZero(t, len(previous), "previous should not be empty")

	assert.Equal(t, buildExpectedDistances(), buildDistToVerify(distances), "expected distances should match to the expected")
	assert.Equal(t, buildExpectedPrevious(), buildPrevToVerify(previous), "expected previous should be match the expected")
}

func TestShortestPathBinary(t *testing.T) {
	testShortestPath(t, heap.NewBinaryMinHeap())
}

func TestShortestPathBinomial(t *testing.T) {
	testShortestPath(t, heap.NewBinomialMinHeap())
}
