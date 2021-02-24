package main

import (
	"log"

	"github.com/marksartdev/assignments/pkg/btree"
)

// nolint:gomnd // It's temporary test.
func main() {
	list1 := &btree.Node{Data: 3}
	list2 := &btree.Node{Data: 13}
	list3 := &btree.Node{Data: 7}
	list4 := &btree.Node{Data: 11}
	list5 := &btree.Node{Data: 6}
	list6 := &btree.Node{Data: 10}
	list7 := &btree.Node{Data: 1}
	list8 := &btree.Node{Data: 15}
	list9 := &btree.Node{Data: 0, Left: list1, Right: list2}
	list10 := &btree.Node{Data: 5, Left: list3, Right: list4}
	list11 := &btree.Node{Data: 4, Right: list5}
	list12 := &btree.Node{Data: 2, Left: list6, Right: list7}
	list13 := &btree.Node{Data: 8, Left: list8, Right: list9}
	list14 := &btree.Node{Data: 10, Left: list10, Right: list11}
	list15 := &btree.Node{Data: 11, Left: list12, Right: list13}
	root := &btree.Node{Data: 7, Left: list14, Right: list15}

	res := btree.BFSearch(root, func(n *btree.Node) bool {
		return root.Data+n.Data == 20
	})

	log.Print(root.Data, res.Data)
}
