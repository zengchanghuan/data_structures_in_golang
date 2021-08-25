package main

import (
	"fmt"
	"io"
	"os"
)

type node struct {
	left *node
	right *node
	data int
}

type BinaryTree struct {
	root *node
}

func (t *BinaryTree)Insert(data int) *BinaryTree  {
	if t.root == nil {
		t.root = &node{data: data,left: nil,right: nil}
	} else {
		t.root.Insert(data)
	}
	return t
}
func (n *node)Insert(data int)  {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data,left: nil,right: nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: data,left: nil,right: nil}
		} else {
			n.right.Insert(data)
		}
	}
}

func print(w io.Writer, node *node, ns int, ch rune) {
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
func main()  {
	tree := &BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}
