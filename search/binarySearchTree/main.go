package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
}

func (n *Node) Print()  {
	fmt.Printf("%d",n.value)
}
func (n *Node) Insert (data int)  {
	if n.value < data {
		if n.right == nil {
			n.right = &Node{value: data}
		} else {
			n.right.Insert(data)
		}
	} else if n.value > data {
		if n.left == nil {
			n.left = &Node{value: data}
		} else {
			n.left.Insert(data)
		}
	}
}

func (n *Node) Search(data int) bool  {
	if n == nil {
		return false
	}
	if n.value < data {
		return n.right.Search(data)
	} else if n.value > data {
		return n.left.Search(data)
	}
	return true
}

func (n *Node) IsEmpty() bool {
	return n.value == 0 && n.left == nil && n.right == nil
}
func (n *Node) FindMin()int  {
	if n.IsEmpty() {
		return 0
	}
	left := (*n).left
	for left.left != nil {
		left = left.left
	}
	return left.value
}

func (n *Node) FindMax() int  {
	if n.IsEmpty() {
		return 0
	}

	for n.right != nil {
		return n.right.FindMax()
	}
	return n.value
}

func (n *Node) PreOrder()  {
	if n.IsEmpty() {
		return
	}
	fmt.Printf("%v ",n.value)

	if n.left != nil {
		n.left.PreOrder()
	}

	if n.right != nil {
		n.right.PreOrder()
	}
}

func (n *Node) MidOrder()  {
	if n.IsEmpty() {
		return
	}

	if n.left != nil {
		n.left.MidOrder()
	}

	fmt.Printf("%v ",n.value)

	if n.right != nil {
		n.right.MidOrder()
	}
}

func (n *Node) PostOrder()  {
	if n.IsEmpty() {
		return
	}

	if n.left != nil {
		n.left.PostOrder()
	}

	if n.right != nil {
		n.right.PostOrder()
	}

	fmt.Printf("%v ",n.value)
}


//层次遍历（广度优先遍历）
func (node *Node) BreadthFirstSearch()  {
	if node == nil {
		return
	}
	result := []int{}
	nodes := []*Node{node}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result,curNode.value)
		if curNode.left != nil {
			nodes = append(nodes,curNode.left)
		}
		if curNode.right != nil {
			nodes = append(nodes,curNode.right)
		}
	}

	for _,v := range result {
		fmt.Printf("%v ",v)
	}
}

//层数（递归实现）深度=左右子树深度的最大值+1
func (node *Node) Layers() int  {
	if node == nil {
		return 0
	}

	leftLayers := node.left.Layers()
	rightLayers := node.right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

func (node *Node) LayersByQueue() int  {
	if node == nil {
		return 0
	}
	layers := 0
	nodes := []*Node{node}
	for len(nodes) > 0 {
		layers++
		size := len(nodes)
		count := 0
		for count < size {
			count++
			curNode := nodes[0]
			nodes = nodes[1:]
			if curNode.left != nil {
				nodes = append(nodes,curNode.left)
			}
			if curNode.right != nil {
				nodes = append(nodes,curNode.right)
			}
		}
	}
	return layers
}

func (node *Node)DeleteNode(data int)  {
	remove(node,data)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.value {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.value {
		node.right = remove(node.right, key)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.value, node.value = leftmostrightside.value, leftmostrightside.value
	node.right = remove(node.right, node.value)
	return node
}
func main()  {
	tree := &Node{value: 8}
	//fmt.Println(tree)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(9)

	tree.PreOrder()
	fmt.Printf("\n")

	tree.DeleteNode(4)
	tree.PreOrder()
	fmt.Printf("\n")


	/*
	fmt.Println(tree.Search(24))
	fmt.Println(tree.FindMin())
	fmt.Println(tree.FindMax())

	tree.PreOrder()
	fmt.Printf("\n")
	tree.MidOrder()
	fmt.Printf("\n")
	tree.PostOrder()


	tree.BreadthFirstSearch()

	fmt.Println("层数为",tree.Layers())
	fmt.Println("层数为",tree.LayersByQueue())
	*/


}
