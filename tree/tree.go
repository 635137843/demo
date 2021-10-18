package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) traverse() {
	//if node == nil{
	//	return
	//}
	//node.left.traverse()
	//node.print()
	//node.right.traverse()

	node.traverseFunc(func(node *treeNode) {
		node.print()
	})
	fmt.Println()
}

func (node *treeNode) traverseFunc(f func(node *treeNode)) {
	if node == nil {
		return
	}
	node.left.traverseFunc(f)
	f(node)
	node.right.traverseFunc(f)
}

func main() {
	var root treeNode
	var left treeNode
	left.value = 0
	root = treeNode{value: 3}
	root.left = &left
	root.right = &treeNode{value: 5}
	root.right.left = new(treeNode)
	root.right.left.setValue(4)
	root.left.right = createNode(2)
	root.traverse()

	nodeCount := 0
	root.traverseFunc(func(node *treeNode) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}
