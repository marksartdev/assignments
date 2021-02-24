package btree

type queue struct {
	queue []*Node
}

func (q *queue) push(n *Node) {
	q.queue = append(q.queue, n)
}

func (q *queue) hasNext() bool {
	return len(q.queue) > 0
}

func (q *queue) next() *Node {
	n := q.queue[0]
	q.queue = q.queue[1:]

	return n
}
