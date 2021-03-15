package test

import (
	"fmt"
	"testing"
	"vb_graph/graph"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	g, err := graph.ParseDIMACS("data/dummy.dimacs", "e")
	assert.Nil(t, err, "ParseDIMACS should not return any error")
	assert.Equal(t, len(g.Edges()), 3, "Edge number should match")
	assert.Equal(t, len(g.Nodes()), 4, "Node number should match")

	graph.Visit(g, graph.DFS, func(n graph.Node) { fmt.Println(n.GetLabel()) }, func(edge graph.Edge) { fmt.Println(edge.GetWeight()) })
	graph.Visit(g, graph.BFS, func(n graph.Node) { fmt.Println(n.GetLabel()) }, func(edge graph.Edge) { fmt.Println(edge.GetWeight()) })
}
