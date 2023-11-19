package main

import (
	"Algorithms/tree/RedBlackTree"
)

func main() {
	tree := RedBlackTree.RedBlackTree{}

	node1 := &RedBlackTree.Node{Key: 6}
	node2 := &RedBlackTree.Node{Key: 7}
	node3 := &RedBlackTree.Node{Key: 10}
	node4 := &RedBlackTree.Node{Key: 15}
	node5 := &RedBlackTree.Node{Key: 17}
	node6 := &RedBlackTree.Node{Key: 19}
	node7 := &RedBlackTree.Node{Key: 1}

	tree.Insert(node1)
	tree.Insert(node2)
	tree.Insert(node3)
	tree.Insert(node4)
	tree.Insert(node5)
	tree.Insert(node6)
	tree.Insert(node7)

	tree.Print()
}
