package dataStructure

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) InsertRoot(data int) *BinaryTree {

	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.Insert(data)
	}
	return t
}

func (n *BinaryNode) Insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.Insert(data)
		}
	}
}



func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func PrintTree() {
	tree := &BinaryTree{}
	tree.InsertRoot(100).
		InsertRoot(-20).
		InsertRoot(-50).
		InsertRoot(-15).
		InsertRoot(-60).
		InsertRoot(50).
		InsertRoot(60).
		InsertRoot(55).
		InsertRoot(85).
		InsertRoot(15).
		InsertRoot(5).
		InsertRoot(-10)
	print(os.Stdout, tree.root, 0, 'M')
}
