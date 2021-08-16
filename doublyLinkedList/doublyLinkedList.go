package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type doublyLinkedList struct {
	tail *node
	head *node
	length int
}

func initDoublyList() *doublyLinkedList  {
	return &doublyLinkedList{}
}

func (d *doublyLinkedList) AddFrontNode(value int)  {
	newData := &node{
		data: value,
	}
	if d.head == nil {
		d.head = newData
		d.tail = newData
	} else {
		newData.next = d.head
		d.head.prev = newData
		d.head = newData
	}
	d.length++
	return
}

func (d *doublyLinkedList) AddEndNode(value int)  {
	newNode := &node{
		data: value,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.length++
	return
}
func (d *doublyLinkedList)TraverseForward()error  {
	if d.head == nil {
		return fmt.Errorf("list is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.data, temp.next)
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.data, temp.next)

		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}
func (d *doublyLinkedList)Size()int  {
	return d.length
}
func main()  {
	doublyList := initDoublyList()
	doublyList.AddEndNode(1)
	doublyList.AddEndNode(2)
	doublyList.AddEndNode(3)
	doublyList.AddEndNode(4)
	doublyList.AddEndNode(5)
	doublyList.AddEndNode(6)
	doublyList.AddEndNode(7)
	doublyList.TraverseForward()
	doublyList.TraverseReverse()
	doublyList.TraverseForward()

}
