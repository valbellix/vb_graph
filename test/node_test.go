package test

import (
	"testing"
	"vb_graph/graph"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	g, err := graph.ParseDIMACS("data/dummy.dimacs", "e")
	assert.Nil(t, err, "ParseDIMACS should not return any error")
	assert.Equal(t, len(g.Edges()), 3, "Edge number should match")
	assert.Equal(t, len(g.Nodes()), 4, "Node number should match")
}
