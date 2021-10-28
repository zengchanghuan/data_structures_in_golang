package main

import (
	"fmt"
)

//循环链表
type Ring struct {
	data       int
	next, prev *Ring
}

func (r *Ring) Init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.Init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.Init()
	}
	return r.prev
}

func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.Init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

//To node a, link a node and return to the back drive node of previous node a
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

//delete n nodes after the node
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

func (r *Ring) Len() int  {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
func deleteTest() {
	//First node
	r := &Ring{data: 9}

	//Link new five nodes
	r.Link(&Ring{data: 8})
	r.Link(&Ring{data: 7})
	r.Link(&Ring{data: 6})
	r.Link(&Ring{data: 5})

	temp := r.Unlink(3) // the last two nodes are removed

	//Print original node
	node := r
	for {
		//Print node values
		fmt.Println(node.data)
		//Move to next node
		node = node.Next()

		//If the node returns to the starting point, the end
		if node == r {
			break
		}
	}

	fmt.Println("------")

	//Print cut-off nodes
	node = temp
	for {
		//Print node values
		fmt.Println(node.data)
		//Move to next node
		node = node.Next()

		//If the node returns to the starting point, the end
		if node == temp {
			break
		}
	}
}
func linkNewTest() {
	r := &Ring{data: -1}

	//Link new five nodes
	r.Link(&Ring{data: 1})
	r.Link(&Ring{data: 2})
	r.Link(&Ring{data: 3})
	r.Link(&Ring{data: 4})

	node := r
	for {
		fmt.Println(node.data)
		node = node.Next()

		if node == r {
			return
		}
	}
}

func main() {
	//linkNewTest()
	deleteTest()
}
