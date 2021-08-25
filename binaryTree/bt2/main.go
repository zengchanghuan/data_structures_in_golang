package main

import "fmt"

type BinarySearchTree struct {
	value int
	left *BinarySearchTree
	right *BinarySearchTree
}

func NewBinarySearchTree() *BinarySearchTree  {
	t := &BinarySearchTree{
		value: 0,
		left: nil,
		right: nil,
	}
	return t
}
func (t *BinarySearchTree) MakeEmpty()  {
	t.value = 0
	t.left = nil
	t.right = nil
}
func (t *BinarySearchTree) IsEmpty()bool  {
	return t.value == 0 && t.left == nil && t.right == nil
}
func (t *BinarySearchTree) Insert(data int) *BinarySearchTree  {
	if t.value == 0 && t.left == nil && t.right == nil {
		t.value = data
		return nil
	}

	if t.value == data {
		return nil
	} else if data < t.value {
		if t.left == nil {
			t.left = &BinarySearchTree{
				value: data,
				left: nil,
				right: nil,
			}
			return nil
		} else {
			t.left.Insert(data)
		}
	} else {
		if data > t.value {
			if t.right == nil {
				t.right = &BinarySearchTree{
					value: data,
					left: nil,
					right: nil,
				}
				return nil
			} else {
				t.right.Insert(data)
			}
		}
	}
	return nil
}
//先根遍历
func (t *BinarySearchTree) PreOrderTraversalRecursion() {
	if (t.IsEmpty()) {
		return
	}
	fmt.Printf("%v \n",t.value)
	if t.left != nil {
		t.left.PreOrderTraversalRecursion()
	}
	if t.right != nil {
		t.right.PreOrderTraversalRecursion()
	}
}
func main()  {
	t := NewBinarySearchTree()
	t.Insert(1)
	t.Insert(89)
	t.Insert(33)
	t.Insert(534)
	t.Insert(313)
	t.Insert(45)
	t.Insert(5353)
	t.PreOrderTraversalRecursion()
}
