package main

import "fmt"

/*
MaxHeap 是一棵完全二叉树，其中父节点的值大于或等于其左右子节点的值。完全二叉树是二叉树，除最后一层外，所有层都是满的。

我们使用一个数组来表示一个 maxheap。根元素是 arr[0]。对于索引 i 我们有

左孩子 – 2*i + 1
右孩子 – 2*i + 2
https://golangbyexample.com/maxheap-in-golang/
*/
type maxHeap struct {
	heapArray []int
	size      int
	maxSize   int
}

func newMaxHeap(maxSize int) *maxHeap {
	maxHeap := &maxHeap{
		heapArray: []int{},
		size:      0,
		maxSize:   maxSize,
	}
	return maxHeap
}

func (m *maxHeap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *maxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *maxHeap) leftchild(index int) int {
	return 2*index + 1
}

func (m *maxHeap) rightchild(index int) int {
	return 2*index + 2
}

func (m *maxHeap) insert(item int) error {
	if m.size >= m.maxSize {
		return fmt.Errorf("Heal is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}
func (m *maxHeap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}
func (m *maxHeap) upHeapify(index int) {
	for m.heapArray[index] > m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *maxHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)

	if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
		largest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex] > m.heapArray[largest] {
		largest = rightRightIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}

func (m *maxHeap) buildMaxHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *maxHeap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}
func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	maxHeap := newMaxHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		maxHeap.insert(inputArray[i])
	}
	maxHeap.buildMaxHeap()
	fmt.Println("The Max Heap is ")
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(maxHeap.remove())
	}
}
