package graph

type VisitType int

const (
	BFS = iota
	DFS
)

func Visit(g Graph, visitType VisitType, visitNode func(Node), visitEdge func(Edge)) {
	nodes := newDequeNode(g.GetRoot())

	l := len(nodes)
	for l != 0 {
		// this is needed to override the outer 'nodes'
		var current Node

		nodes, current = pop(nodes)
		visitNode(current)

		current.MarkVisited()

		for _, edge := range current.OutEdges() {
			visitEdge(edge)

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
