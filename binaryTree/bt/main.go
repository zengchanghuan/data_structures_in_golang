package main

import "fmt"

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
var count int
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main()  {
	tree := Node{Key: 100}
	tree.Insert(200)
	tree.Insert(50)
	tree.Insert(300)
	tree.Insert(4)
	fmt.Println(tree.Search(33))
	fmt.Println(tree.Search(50))
	fmt.Println(count)

}
