package graph

type VisitType int

const (
	BFS = iota
	DFS
)

type Visitor interface {
	VisitNode(Node)
	VisitEdge(Edge)
}

func Visit(g Graph, visitType VisitType, visitor Visitor) {
	g.ResetVisited()

	nodes := newDequeNode(g.GetRoot())

	l := len(nodes)
	for l != 0 {
		// this is needed to override the outer 'nodes'
		var current Node

		nodes, current = pop(nodes)
		visitor.VisitNode(current)

		current.MarkVisited()

		for _, edge := range current.OutEdges() {
			visitor.VisitEdge(edge)

			if !edge.GetDst().IsVisited() {
				if visitType == BFS {
					nodes = append(nodes, edge.GetDst())
				} else {
					nodes = push(nodes, edge.GetDst())
				}
			}
		}
		l = len(nodes)
	}
}
