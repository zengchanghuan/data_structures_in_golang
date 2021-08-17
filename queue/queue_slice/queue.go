package main

import "fmt"

type Queue struct {
	slice []int
}

//入队
func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice,i)
}

//出队
func (q *Queue) Dequeue() int  {
	//取第一个元素
	ret := q.slice[0]

	//remove the first item from the queue
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

func (q *Queue)Peek()int  {
	if q.Empty() {
		return -1
	}
	return q.slice[0]
}
func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func (q *Queue)Size()int  {
	return len(q.slice)
}

func (q *Queue) Empty()bool  {
	return len(q.slice)  == 0
}
func (q *Queue)Clear()  {
	q.slice = nil
}
func main()  {
	var q *Queue = new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Empty())
	q.Clear()
	fmt.Println(q.Empty())
}


