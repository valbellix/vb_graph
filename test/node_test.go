package test

import (
	"testing"
	"vb_graph/graph"

	"github.com/stretchr/testify/assert"
)

type testVisitor struct {
	t *testing.T
}

func (v *testVisitor) VisitEdge(edge graph.Edge) {
	v.t.Logf("Edge with weight: %v\n", edge.GetWeight())
}

func (v *testVisitor) VisitNode(node graph.Node) {
	v.t.Logf("Node labelled: %v\n", node.GetLabel())
}

func TestNode(t *testing.T) {
	g, err := graph.ParseDIMACS("data/dummy.dimacs", "e")
	assert.Nil(t, err, "ParseDIMACS should not return any error")
	assert.Equal(t, len(g.Edges()), 3, "Edge number should match")
	assert.Equal(t, len(g.Nodes()), 4, "Node number should match")

	v := &testVisitor{t}

	t.Log("Testing DFS visit")
	graph.Visit(g, graph.DFS, v)
	t.Log("Testing BFS visit")
	graph.Visit(g, graph.BFS, v)
}
