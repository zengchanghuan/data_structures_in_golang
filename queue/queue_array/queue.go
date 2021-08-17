package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue)Enqueue(i int)  {
	q.items = append(q.items,i)
}

func (q *Queue)Empty()bool  {
	return len(q.items)  == 0
}

func (q *Queue)Size()int  {
	return len(q.items)
}
func (q *Queue)Peek()int  {
	if q.Empty() {
		return -1
	}
	return q.items[0]
}

func (q *Queue) Clear()  {
	q.items = nil
}
func (q *Queue) Dequeue() int  {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
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
