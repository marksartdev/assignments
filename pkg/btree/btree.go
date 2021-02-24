// Package btree contains structs and functions for working with b-tree.
package btree

// Node of b-tree.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}
