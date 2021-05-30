package dataStructure

import (
	"fmt"
	"log"
)

func PreorderRecursive(root *BinaryNode) {
	if root != nil {
		fmt.Printf("%d ->", root.data)
		PreorderRecursive(root.left)
		PreorderRecursive(root.right)
	}
}

func InorderRecursive(root *BinaryNode) {
	if root != nil {
		InorderRecursive(root.left)
		fmt.Printf("%d ->", root.data)
		InorderRecursive(root.right)
	}
}

func PostorderRecursive(root *BinaryNode) {
	if root != nil {
		PostorderRecursive(root.left)
		PostorderRecursive(root.right)
		fmt.Printf("%d ->", root.data)
	}
}

func getDepth(n *BinaryNode) int {
	if n == nil {
		return 0
	}
	left := getDepth(n.left)
	right := getDepth(n.right)
	if left > right {
		log.Println("left")
		return left + 1
	}
	return right + 1
}

func Runner() {
	// r := &BinaryTree{}
	// r.InsertRoot(10)
	// r.InsertRoot(20)
	// r.InsertRoot(100)
	// r.InsertRoot(120)
	// r.InsertRoot(2)
	root := &BinaryNode{data: 10, left: nil, right: nil}
	root.left = &BinaryNode{data: 20, left: nil, right: nil}
	root.right = &BinaryNode{data: 30, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 40, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 50, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 60, left: nil, right: nil}
	PreorderRecursive(root)
	log.Println("--------------")
	PostorderRecursive(root)
	log.Println("--------------")
	InorderRecursive(root)
}
