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
// FindMin 获取最小的值.
// 采用非递归方式.
func (b *BinarySearchTree) FindMin() int {
	if b.IsEmpty() {
		return 0
	}

	if b.left != nil {
		return b.left.FindMin()
	}
	return b.value
}
func (t *BinarySearchTree) FindMax() int  {
	if t.IsEmpty() {
		return -1
	}
	for t.right != nil {
		return t.right.FindMax()
	}
	return t.value
}

// Delete 删除节点, 根据要删除节点子节点的个数分三种情况：
// 1.要删除的节点是叶子节点，直接设置删除节点的父节点指向删除节点的指针为nil
// 2.要删除的节点只有一个节点，设置删除节点的父节点指向删除节点的指针指向"要删除节点的子节点"
// 3.要删除的节点有两个节点，从删除节点的右子树查询最小节点值替换到要删除节点的值，同时删除最小节点
func (b *BinarySearchTree) Delete(data int) {
	// 删除节点的父节点
	pp := b
	// 待删除的节点
	p := b

	for p.value != 0 && p.value != data {
		pp = p
		if data > p.value {
			p = p.right
		} else {
			p = p.left
		}
	}
	// 没有找到要删除的节点
	if p == nil {
		return
	}

	// 第一种情况, 判断要删除节点是位于左节点还是右节点
	if p.left == nil && p.right == nil {
		if pp.left == p {
			pp.left = nil
		} else {
			pp.right = nil
		}
	}

	// 第二种情况，只有一个节点
	if (p.left == nil && p.right != nil) || (p.left != nil && p.right == nil) {
		if pp.left == p {
			pp.left = p.left
		} else {
			pp.right = p.right
		}
	}

	// 第三种情况, 删除节点有两个子节点
	if p.left != nil && p.right != nil {
		// minP的父节点
		minPP := p
		// 查找子右节点的最小节点
		minP := p.right
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}
		// 把右子树中最小节点值替换到要删除节点的值
		p.value = minP.value

		// 删除minP节点
		// 如果最小节点是左叶子节点, 直接删除最小节点
		if minPP.left == minP {
			minPP.left = nil
			// 最小节点不是左叶子节点, 说明当前minP只包含右子树
		} else {
			minPP.right = minP.right
		}
	}
}

// MidOrderTraversalRecursion 递归中序遍历.
// 先遍历左子树，然后遍历根节点，再遍历右子树.
func (b *BinarySearchTree) MidOrderTraversalRecursion() {
	if b.IsEmpty() {
		return
	}

	if b.left != nil {
		b.left.MidOrderTraversalRecursion()
	}
	fmt.Printf("%v ", b.value)
	if b.right != nil {
		b.right.MidOrderTraversalRecursion()
	}
}
// PostOrderTraversalRecursion 递归后序遍历.
// 先遍历左子树，然后遍历右子树，再遍历根节点.
func (b *BinarySearchTree) PostOrderTraversalRecursion() {
	if b.IsEmpty() {
		return
	}

	if b.left != nil {
		b.left.PostOrderTraversalRecursion()
	}
	if b.right != nil {
		b.right.PostOrderTraversalRecursion()
	}
	fmt.Printf("%v ", b.value)
}
func main()  {
	t := NewBinarySearchTree()
	t.Insert(3)
	t.Insert(89)
	t.Insert(33)
	t.Insert(534)
	t.Insert(313)
	t.Insert(45)
	t.Insert(5353)
	t.Delete(534)
	fmt.Println(t.FindMin())
	//t.PreOrderTraversalRecursion()

}
