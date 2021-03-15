package graph

type dequeNode []Node

func push(d dequeNode, n Node) dequeNode {
	return append(dequeNode{n}, d...)
}

func pop(d dequeNode) (dequeNode, Node) {
	n := d[0]

	d[0] = nil
	return d[1:], n
}

func newDequeNode(n Node) dequeNode {
	return []Node{n}
}
