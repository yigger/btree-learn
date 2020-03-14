package main

import (
	"github.com/yigger/btree/tree"
)

func main() {
	t := tree.CreateTree(40)
	nodes := []tree.T{20, 10, 30, 5, 80, 60, 100, 50}
	t.AddNodes(nodes)
	t.PreorderTraversal(t.Root)
}
