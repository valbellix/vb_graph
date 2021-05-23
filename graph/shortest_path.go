package graph

import (
	"math"
	"vb_graph/heap"
)

func initDistances(g Graph, n Node) map[Node]int {
	nodes := g.Nodes()
	m := make(map[Node]int, len(nodes))
	for _, current := range nodes {
		m[current] = math.MaxInt64
	}

	return m
}

func initPrevious(g Graph) map[Node]Node {
	nodes := g.Nodes()
	m := make(map[Node]Node, len(nodes))
	for _, current := range nodes {
		m[current] = nil
	}

	return m
}

type nodeHolder struct {
	node     Node
	distance int
}

func (n *nodeHolder) Priority() int {
	return n.distance
}

func (n *nodeHolder) SetPriority(p int) {
	n.distance = p
}

func ShortestPath(g Graph, n Node) (map[Node]int, map[Node]Node, error) {
	distance := initDistances(g, n)
	previous := initPrevious(g)

	h := heap.NewBinaryHeap(heap.MIN_HEAP)
	h.Push(&nodeHolder{n, 0})
	distance[n] = 0

	for !h.IsEmpty() {
		currentNode := h.Pop().(*nodeHolder).node
		currentNodeDistance := distance[currentNode]

		// exploring neighbours of currentNode
		for _, edge := range currentNode.OutEdges() {
			dest := edge.GetDst()
			w := edge.GetWeight()
			currentDstDistance := distance[dest]

			tmpDist := currentNodeDistance + w
			if tmpDist < currentDstDistance {
				distance[dest] = tmpDist
				previous[dest] = currentNode

				h.Push(&nodeHolder{dest, tmpDist})
			}
		}
	}

	return distance, previous, nil
}
