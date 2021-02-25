package main

import (
	"fmt"

	"github.com/marksartdev/assignments/pkg/btree"
)

// nolint:gomnd // It's temporary test.
func main() {
	k := 7

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

	pairs := make(map[[2]int]struct{})

	pairCh := make(chan [2]int, 10)
	exitCh := make(chan struct{}, 10)
	goroutineCount := 0

	var iterator btree.Iterator = root.GetIterator()
	for iterator.HasNext() {
		next := iterator.Next()
		startGoroutine(pairCh, exitCh, root, next, k)
		goroutineCount++
	}

	for goroutineCount > 0 {
		select {
		case pair := <-pairCh:
			pairs[pair] = struct{}{}
			goroutineCount--
		case <-exitCh:
			goroutineCount--
		}
	}

	for pair := range pairs {
		fmt.Println(pair) // nolint:forbidigo // It's for printing results.
	}
}

func startGoroutine(pairCh chan [2]int, exitCh chan struct{}, root, base *btree.Node, k int) {
	go func() {
		if res := bfSearch(root, getCheckFunc(base, k)); res != nil {
			sendPair(pairCh, base.Data, res.Data)
		} else {
			exitCh <- struct{}{}
		}
	}()
}

func getCheckFunc(base *btree.Node, k int) func(node *btree.Node) bool {
	return func(node *btree.Node) bool {
		if node == base {
			return false
		}

		return base.Data+node.Data == k
	}
}

func sendPair(pairCh chan [2]int, a, b int) {
	if a < b {
		pairCh <- [2]int{a, b}
	} else {
		pairCh <- [2]int{b, a}
	}
}

func bfSearch(root *btree.Node, check func(node *btree.Node) bool) *btree.Node {
	iterator := root.GetIterator()
	for iterator.HasNext() {
		node := iterator.Next()

		if check(node) {
			return node
		}
	}

	return nil
}
