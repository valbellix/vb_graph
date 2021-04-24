package graph

import (
	"math"
	"vb_graph/heap"
)

func initDistances(g Graph, n Node) map[Node]int {
	m := make(map[Node]int)
	for _, current := range g.Nodes() {
		m[current] = math.MaxInt64
	}

	return m
}

func initPrevious(g Graph) map[Node]Node {
	m := make(map[Node]Node)
	for _, current := range g.Nodes() {
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

	g.ResetVisited()

	h := heap.NewBinaryHeap(heap.MIN_HEAP)
	h.Push(&nodeHolder{n, 0})
	distance[n] = 0

	for !h.IsEmpty() {
		currentNode := h.Pop().(*nodeHolder).node
		for _, edge := range currentNode.OutEdges() {
			dest := edge.GetDst()
			w := edge.GetWeight()
			currentDstDistance := distance[dest]
			if currentDstDistance == math.MaxInt64 {
				h.Push(&nodeHolder{dest, w})
			} else if (distance[currentNode] + w) < currentDstDistance {
				newDist := distance[currentNode] + w
				distance[dest] = newDist
				previous[currentNode] = dest
				err := h.MoveUp(&nodeHolder{dest, currentDstDistance}, newDist)
				if err != nil {
					return nil, nil, err
				}
			}
		}
	}

	return distance, previous, nil
}
