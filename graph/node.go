package graph

type Edge interface {
	GetSrc() Node
	GetDst() Node

	GetWeight() int
}

type Node interface {
	GetLabel() string
	OutEdges() []Edge

	AddEdge(e Edge)
	RemoveEdge(e Edge)

	MarkVisited()
	ResetVisited()
	IsVisited() bool
}

type adjNode struct {
	outEdges []Edge
	label    string
	visited  bool
}

func NewNode(l string) Node {
	return &adjNode{
		label:    l,
		outEdges: make([]Edge, 0),
		visited:  false,
	}
}

func (n *adjNode) GetLabel() string {
	return n.label
}

func (n *adjNode) OutEdges() []Edge {
	return n.outEdges
}

func (n *adjNode) AddEdge(e Edge) {
	n.outEdges = append(n.outEdges, e)
}

func (n *adjNode) RemoveEdge(e Edge) {
	// TODO
}

func (n *adjNode) IsVisited() bool {
	return n.visited
}

func (n *adjNode) MarkVisited() {
	n.visited = true
}

func (n *adjNode) ResetVisited() {
	n.visited = false
}

type adjEdge struct {
	src Node
	dst Node

	weight int
}

func NewEdge(src, dst Node, weight int) Edge {
	return &adjEdge{
		src:    src,
		dst:    dst,
		weight: weight,
	}
}

func (e *adjEdge) GetSrc() Node {
	return e.src
}

func (e *adjEdge) GetDst() Node {
	return e.dst
}

func (e *adjEdge) GetWeight() int {
	return e.weight
}
