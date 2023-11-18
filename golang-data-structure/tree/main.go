package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// 前序遍历
func PreOrder(node *Node) {
	if node != nil {
		fmt.Print(" ", node.Data)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历
func InOrder(node *Node) {
	if node != nil {
		InOrder(node.Left)
		fmt.Print(" ", node.Data)
		InOrder(node.Right)
	}
}

// 后序遍历
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Print(" ", node.Data)
	}
}

func CreateTreeNode() *Node { //p80
	node9 := &Node{
		Data:  9,
		Left:  nil,
		Right: nil,
	}

	node10 := &Node{
		Data:  10,
		Left:  nil,
		Right: nil,
	}

	node4 := &Node{
		Data:  4,
		Left:  nil,
		Right: nil,
	}

	node2 := &Node{
		Data:  2,
		Left:  node9,
		Right: node10,
	}

	node8 := &Node{
		Data:  8,
		Left:  nil,
		Right: node4,
	}

	node3 := &Node{
		Data:  3,
		Left:  node2,
		Right: node8,
	}
	return node3
}

func main() {
	rootNode := CreateTreeNode()
	fmt.Println("先序遍历")
	PreOrder(rootNode)
	fmt.Println()
	fmt.Println("中序遍历")
	InOrder(rootNode)
	fmt.Println()
	fmt.Println("后序遍历")
	PostOrder(rootNode)
}
