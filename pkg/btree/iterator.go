package btree

// Iterator for B-Tree.
type Iterator interface {
	HasNext() bool
	Next() *Node
}

type treeIterator struct {
	queue []*Node
}

// HasNext element in Iterator.
func (q *treeIterator) HasNext() bool {
	return len(q.queue) > 0
}

// Next element in Iterator.
func (q *treeIterator) Next() *Node {
	n := q.queue[0]
	q.queue = q.queue[1:]

	if n.Left != nil {
		q.queue = append(q.queue, n.Left)
	}

	if n.Right != nil {
		q.queue = append(q.queue, n.Right)
	}

	return n
}
