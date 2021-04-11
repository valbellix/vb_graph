package graph

import (
	"math"
	"vb_graph/heap"
)

func initMap(g Graph, n Node) map[Node]int {
	m := make(map[Node]int)
	for _, current := range g.Nodes() {
		m[current] = math.MaxInt64
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

func ShortestPath(g Graph, n Node) (map[Node]int, error) {
	result := initMap(g, n)
	g.ResetVisited()

	h := heap.NewBinaryHeap(heap.MIN_HEAP)
	h.Push(&nodeHolder{n, 0})
	result[n] = 0

	for !h.IsEmpty() {
		currentNode := h.Pop().(*nodeHolder).node
		for _, edge := range currentNode.OutEdges() {
			dest := edge.GetDst()
			w := edge.GetWeight()
			currentDstDistance := result[dest]
			if currentDstDistance == math.MaxInt64 {
				h.Push(&nodeHolder{dest, w})
			} else if (result[currentNode] + w) < currentDstDistance {
				newDist := result[currentNode] + w
				result[dest] = newDist
				err := h.MoveUp(&nodeHolder{dest, currentDstDistance}, newDist)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return result, nil
}
