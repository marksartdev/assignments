package btree

// BFSearch Node.
func BFSearch(node *Node, check func(*Node) bool) *Node {
	q := queue{}
	q.push(node)

	for q.hasNext() {
		n := q.next()

		if check(n) {
			return n
		}

		if n.Left != nil {
			q.push(n.Left)
		}

		if n.Right != nil {
			q.push(n.Right)
		}
	}

	return nil
}
