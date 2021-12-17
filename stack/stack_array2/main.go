package main

import "fmt"

type Stack struct {
	slice []int
}

func (s *Stack) Push(i int)  {
	s.slice = append(s.slice,i)
}

func (s *Stack) Peek() int  {
	return s.slice[len(s.slice) - 1]
}

func (s *Stack) Pop() int  {
	var  ret int = s.Peek()
	s.slice = s.slice[0:len(s.slice) - 1]
	return ret
}

func (s *Stack) String() string  {
	return fmt.Sprint(s.slice)
}
func main()  {
	var s *Stack = new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)
	s.Push(99)
	fmt.Println(s)
}