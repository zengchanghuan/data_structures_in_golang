package main

import "fmt"

/*
MinHeap 是一棵完全二叉树，其中父节点的值小于或等于其左右子节点的值。完全二叉树是二叉树，除最后一层外，所有层都是满的。
左孩子 – 2*i + 1
右孩子 – 2*i + 2

*/
type minHeap struct {
	heapArray []int
	size      int
	maxSize   int
}

func newMinHeap(maxsize int) *minHeap {
	minHeap := &minHeap{
		heapArray: []int{},
		size:      0,
		maxSize:   maxsize,
	}
	return minHeap
}

func (m *minHeap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *minHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *minHeap) leftChild(index int) int {
	return 2*index + 1
}
func (m *minHeap) rightChild(index int) int {
	return 2*index + 2
}

func (m *minHeap) insert(item int) error {
	if m.size >= m.maxSize {
		return fmt.Errorf("heal is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *minHeap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}
func (m *minHeap) upHeapify(index int) {
	for m.heapArray[index] < m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *minHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}

	smallest := current
	leftChildIndex := m.leftChild(current)
	rightRightIndex := m.rightChild(current)

	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
		smallest = leftChildIndex
	}

	if rightRightIndex < m.size && m.heapArray[rightRightIndex] < m.heapArray[smallest] {
		smallest = rightRightIndex
	}

	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
	return

}

func (m *minHeap) buildMinHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}
func (m *minHeap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}
func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := newMinHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		minHeap.insert(inputArray[i])
	}
	minHeap.buildMinHeap()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.remove())
	}
	fmt.Scanln()

}
