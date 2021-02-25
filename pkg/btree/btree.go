// Package btree contains structs and functions for working with b-tree.
package btree

// Node of b-tree.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// GetIterator for B-Tree.
func (n *Node) GetIterator() Iterator {
	return &treeIterator{queue: []*Node{n}}
}
