package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (s *Stack) Peek() int {
	data := s.items[0]
	return data
}
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)

}
