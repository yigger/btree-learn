package tree

import "fmt"

type T int

type Node struct {
	LeftChild  *Node
	RightChild *Node
	Val        T
}

type Tree struct {
	Root *Node
}

var tree *Tree

func CreateTree(value T) *Tree {
	if tree == nil {
		tree = &Tree{
			Root: &Node{
				LeftChild:  nil,
				RightChild: nil,
				Val:        value,
			},
		}
	}
	return tree
}
func (t *Tree) AddNodes(values []T) {
	for _, item := range values {
		t.AddNode(item)
	}
}

func (t *Tree) AddNode(value T) {
	node := traver(t.Root, value)
	newNode := &Node{
		LeftChild:  nil,
		RightChild: nil,
		Val:        value,
	}
	if node.LeftChild == nil {
		node.LeftChild = newNode
	} else {
		node.RightChild = newNode
	}
}

func traver(node *Node, val T) *Node {
	if (node.LeftChild == nil && node.Val > val) || (node.RightChild == nil && node.Val < val) {
		return node
	}

	if node.Val > val {
		return traver(node.LeftChild, val)
	} else {
		return traver(node.RightChild, val)
	}
}

// 先序遍历
func (t *Tree) PreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	t.PreorderTraversal(node.LeftChild)
	t.PreorderTraversal(node.RightChild)
}
