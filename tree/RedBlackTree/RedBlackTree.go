package RedBlackTree

import "fmt"

type RedBlackTree struct {
	root *Node
}

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
	Color  Color
}
type Color bool

const (
	Black Color = true
	Red         = false
)

const (
	colorRed  = "\033[31m"
	colorBlue = "\033[34m"
)

func (tree *RedBlackTree) Insert(node *Node) {

	if node == nil {
		return
	}

	if tree.root == nil {
		tree.root = node
		tree.root.Color = Black
		return
	}

	node.Color = Red
	current := tree.root
	var parent *Node

	for current != nil {
		if node.Key < current.Key {
			parent = current
			current = current.Left
		} else {
			parent = current
			current = current.Right
		}
	}

	if node.Key < parent.Key {
		parent.Left = node
		node.Parent = parent
	} else {
		parent.Right = node
		node.Parent = parent
	}

	if node.Parent != nil && node.Parent.Color == Red {
		tree.leftBalance(node)
		tree.root.Color = Black
	}

}

func (tree *RedBlackTree) fixTree(node *Node) {

	if node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			tree.rightBalance(node)
		} else {
			tree.leftBalance(node)

		}
	}
	tree.root.Color = Black
}

func (tree *RedBlackTree) rightBalance(node *Node) {
	uncle := node.Parent.Parent.Right

	if uncle != nil && uncle.Color == Red {
		node.Parent.Color = Black
		uncle.Color = Black
	} else {
		tree.rightRotate(node)
	}
}

func (tree *RedBlackTree) leftBalance(node *Node) {
	uncle := node.Parent.Parent.Left

	if uncle != nil && uncle.Color == Red {
		node.Parent.Color = Black
		uncle.Color = Black
	} else {
		tree.leftRotate(node)
	}
}

func (tree *RedBlackTree) leftRotate(node *Node) {
	parent := node.Parent
	grandParent := parent.Parent
	greatGrandParent := grandParent.Parent

	parent.Left = grandParent
	parent.Parent = greatGrandParent
	grandParent.Parent = parent
	grandParent.Right = nil

	if greatGrandParent == nil {
		tree.root = parent
	}
	grandParent.Color = Red
}

func (tree *RedBlackTree) rightRotate(node *Node) {
	parent := node.Parent
	grandParent := parent.Parent
	greatGrandParent := grandParent.Parent

	parent.Right = grandParent
	parent.Parent = greatGrandParent
	grandParent.Parent = parent
	grandParent.Left = nil
	if greatGrandParent == nil {
		tree.root = parent
	}
	grandParent.Color = Red
}

func (tree *RedBlackTree) GetHeight() (height int) {
	if tree.root == nil {
		return 0
	}
	cur := tree.root
	for height = 1; cur != nil; height++ {
		cur = cur.Left
	}

	return height
}

func (tree *RedBlackTree) Print() {
	cur := tree.root
	tree.draw(cur)
}

func (tree *RedBlackTree) draw(node *Node) {

	if node != nil {
		if node.Color == Red {
			fmt.Println(colorRed, node.Key)
		} else {
			fmt.Println(colorBlue, node.Key)
		}
		if node.Left != nil || node.Right != nil {
			tree.draw(node.Left)
			tree.draw(node.Right)
		}
	}
}
