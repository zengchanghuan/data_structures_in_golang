package main

type node struct {
	data int
	pre  *node
	next *node
}

type Queue struct {
	head *node //头指针
	tail *node //尾指针
	size int
}

func (q *Queue) Enqueue(e int) {

}
func (q *Queue) Empty() bool {
	return q.head == q.tail
}
func (q *Queue) Dequeue() int {

}

func (q *Queue) Peek() int {
	if q.Empty() {
		return -1
	}
	return q.head.next.data
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Clear() {
	q.head.next = nil
	q.tail = q.head
	q.size = 0
}
func main() {

}
