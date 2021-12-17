package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(i int)  {
	q.slice = append(q.slice,i)
}

func (q *Queue) Dequeue() int {
	ret := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

func (q *Queue) String() string  {
	return fmt.Sprint(q.slice)
}

func main()  {
	var q *Queue = new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}


