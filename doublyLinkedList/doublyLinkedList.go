package main

import "fmt"
/*
双向链表
https://golangbyexample.com/doubly-linked-list-golang/
https://github.com/l3x/golang-code-examples/blob/master/doubly-linked-list.go#L23
*/
type node struct {
	data int
	next *node
	prev *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	length int
}
func (n *node) Next() *node  {
	return n.next
}
func (n *node) Prev() *node  {
	return n.prev
}
func initDoublyList() *doublyLinkedList  {
	return &doublyLinkedList{}
}
func (l *doublyLinkedList) First() *node  {
	return l.head
}

func (l *doublyLinkedList) Last() *node  {
	return l.tail
}

//头插法
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

func (d *doublyLinkedList)Find(value int) *node  {
	found := false
	var ret *node = nil
	for n := d.First();n != nil && !found; n = n.Next() {
		if n.data == value {
			found = true
			ret = n
		}
	}
	return ret
}

func (d *doublyLinkedList)Delete(value int)bool  {
	success := false
	deleteNode := d.Find(value)
	if deleteNode != nil {
		fmt.Println("Delete - FOUND: ", value)
		preNode := deleteNode.prev
		nextNode := deleteNode.next

		//remove this node
		preNode.next = deleteNode.next
		nextNode.prev = deleteNode.prev
		success = true
	}
	return success
}
func main()  {
	doublyList := initDoublyList()
	doublyList.AddFrontNode(1)
	doublyList.AddFrontNode(2)
	doublyList.AddFrontNode(3)
	doublyList.AddFrontNode(4)
	doublyList.AddFrontNode(5)
	doublyList.AddFrontNode(6)
	doublyList.AddFrontNode(7)
	//fmt.Println(doublyList.TraverseForward())
	//ret := doublyList.Find(88)
	//fmt.Println(ret.prev.data,ret.next.data)
	//doublyList.TraverseReverse()
	//doublyList.TraverseForward()

	//fmt.Println(doublyList.Delete(88))
	doublyList.TraverseReverse()

}
