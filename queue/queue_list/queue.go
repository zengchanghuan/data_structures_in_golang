package main

import "fmt"

type node struct {
	data int
	next *node
}

type Queue struct {
	head *node //队头指针
	tail *node //队尾指针
	len  int
}

func InitQueue() *Queue {
	q := &Queue{
		len:  0,
		head: nil,
		tail: nil,
	}
	return q
}

func (q *Queue) Enqueue(data int) {
	node := &node{data: data, next: nil}
	if q.IsEmpty(){
		q.head = node
		q.tail = node
		q.len++
		return
	}

	q.tail.next = node
	q.tail = node
	q.len++
}

func (q *Queue) Dequeue()  {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}
	if q.len == 1 {
		q.head = nil
		q.tail = nil
		q.len--
		return
	}
	q.head = q.head.next
	q.len--
}

func (q *Queue) IsEmpty()bool  {
	if q.len == 0 {
		return true
	}
	return false
}
func (q *Queue) Print()  {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
	}

	p := q.head
	index := 1
	for p != nil {
		fmt.Printf("index == %#v,Data == %#v\n", index, p.data)
		index++
		p = p.next
	}
}

func (q *Queue)Length()int  {
	return q.len
}

func (q *Queue)GetHead()int  {
	if q.head != q.tail {
		return q.head.data
	}
	return -1
}

func (q *Queue)Clear()  {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
	}

	q.head.next = nil
	q.tail.next = nil
	q.len = 0
}


func main() {
	q := InitQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	//fmt.Println(q.Length())
	q.Print()
	//fmt.Println(q.Length())
	//fmt.Println(q.GetHead())
	//q.Dequeue()
	fmt.Println(q.GetHead())
	q.Clear()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Length())
	//q.Dequeue()
	//fmt.Println(q.GetHead())

	//q.Print()
	//fmt.Println(q.Length())

	//q.Dequeue()
	//q.Print()

}
