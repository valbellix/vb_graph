package graph

import "errors"

type Graph interface {
	Edges() []Edge
	Nodes() []Node

	SetRoot(n Node) error
	GetRoot() Node

	AverageOutboundDegree() float32

	AddEdge(from, to string, weight int) error

	GetNode(label string) (Node, bool)

	ResetVisited()
}

type adjGraph struct {
	root  *adjNode
	nodes map[string]Node
	edges []Edge
}

func NewGraph() Graph {
	return &adjGraph{
		root:  nil,
		nodes: make(map[string]Node),
		edges: make([]Edge, 0),
	}
}

func (g *adjGraph) Edges() []Edge {
	return g.edges
}

func (g *adjGraph) Nodes() []Node {
	nodes := make([]Node, 0)

	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}
	return nodes
}

func (g *adjGraph) SetRoot(n Node) error {
	if g.root == nil {
		g.root = n.(*adjNode)
		g.nodes[n.GetLabel()] = n
		return nil
	} else {
		if _, ok := g.nodes[n.GetLabel()]; ok {
			g.root = n.(*adjNode)
			return nil
		}
		return errors.New("if the graph already has a root, it is not possible to set as root a node that does not belong to the graph")
	}
}

func (g *adjGraph) GetRoot() Node {
	return g.root
}

func (g *adjGraph) AverageOutboundDegree() float32 {
	tot := 0
	for _, n := range g.nodes {
		tot += len(n.OutEdges())
	}

	if tot == 0 {
		return 0
	}

	return float32(tot) / float32(len(g.edges))
}

func (g *adjGraph) AddEdge(src, dst string, weight int) error {
	if src == dst {
		return errors.New("it is not possible to add an edge whose source and destination are the same")
	}

	s, nodeInGraph := g.GetNode(src)
	if !nodeInGraph {
		s = NewNode(src)
		if g.root == nil {
			g.root = s.(*adjNode)
		}

		g.nodes[src] = s
	}

	d, nodeInGraph := g.GetNode(dst)
	if !nodeInGraph {
		d = NewNode(dst)
		g.nodes[dst] = d
	}

	edge := NewEdge(s, d, weight)
	s.AddEdge(edge)
	g.edges = append(g.edges, edge)

	return nil
}

func (g *adjGraph) GetNode(label string) (Node, bool) {
	n, ok := g.nodes[label]
	return n, ok
}

func (g *adjGraph) ResetVisited() {
	for _, n := range g.nodes {
		n.ResetVisited()
	}
}
