package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	length int
	head *node
}

func InitLinkStack() *stack  {
	node := new(node)
	l := new(stack)
	l.head = node
	l.length = 0
	return l
}
func (s *stack) Push(value int)  {
	node := &node{
		data: value,
		next: s.head,
	}
	s.head = node
	s.length++
}

func (s *stack) Pop()int  {
	if s.head == nil {
		return -1
	}

	val := s.head.data
	s.head = s.head.next
	s.length--
	return val
}

func (s *stack)Peek()int  {
	if s.head == nil {
		fmt.Println("空栈")
		return -1
	}
	val := s.head.data
	return val
}

func (s *stack)ShowAll()  {
	if s.head.next == nil {
		fmt.Println("stack is empty")
		return
	}

	cur := s.head
	for {
		if cur.next != nil {
			fmt.Printf("%v\n",cur.data)
			cur = cur.next
		} else {
			break
		}
	}
}
func main()  {
	s := InitLinkStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	//s.ShowAll()
	//fmt.Println("*******")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println("*******")
	s.ShowAll()

}