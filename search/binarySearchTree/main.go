package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(value int)  {
	newNode := &Node{value: value,left: nil,right: nil}

	//如果二叉树为空，那么这个节点就当作根节点
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root,newNode)
	}
}

//从根节点依次比较
func insertNode(root, newNode *Node) {
	if newNode.value < root.value { // 应该放到根节点的左边
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value { // 应该放到根节点的右边
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
	// 否则等于根节点
}

func (bst *BinarySearchTree) Search(value int) bool  {
	return search(bst.root,value)
}

func search(n *Node,value int) bool  {
	if n == nil {
		return false
	}
	if value < n.value {
		return search(n.left,value)
	}
	if value > n.value {
		return search(n.right,value)
	}
	return true
}



func main()  {
	tree := new(BinarySearchTree)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(13)
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(66))

}
