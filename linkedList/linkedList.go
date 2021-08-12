package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func initNode(data int) *node {
	return &node{data: data}
}

func initLinkedList() *linkedList {
	return &linkedList{initNode(0), 0}
}

func (l *linkedList) IsEmpty() bool {
	if l.length > 0 {
		return true
	} else {
		return false
	}
}

//清空链表数据，链表仍存在
func (l *linkedList) clear() {
	if l.head == nil {
		return
	}
	l.head.next = nil
	l.length = 0

}

//销毁链表 整张表都不存在
func (l *linkedList) Destroy() {
	if l.head == nil {
		return
	}
	l.length = 0
	l.head = nil

}

//表长
func (l *linkedList) Length() int {
	if l.head == nil {
		return 0
	}
	return l.length
}

//
func (l *linkedList) GetElemWithIndex(index int) *node {
	index -= 1
	if index < 0 {
		fmt.Println("index is error")
		return nil
	}
	if l.length == 0 || index > l.length {
		return nil
	}

	//从链表头指值出发，顺着链域next逐个往下搜索
	cursor := l.head
	var i = 0
	for ; i < index; i++ {
		cursor = cursor.next
	}
	return cursor
}

//查找，按值查找，根据数据获取该数据所在的地址
func (l *linkedList) FindWithData(data int) *node {
	if l.head == nil || l.length == 0 {
		return nil
	}

	current := l.head
	for {
		if current.data == data {
			break
		}
		if current.next == nil {
			break
		}
		current = current.next
	}
	return current

}

//查找，按值查找返回位置序号
func (l *linkedList) LocateElem(data int) int {
	if l.head == nil || l.length == 0 {
		return -1
	}

	cursor := l.head
	index := 1
	//如果没有找到，cursor.next 为空，故下面cursor不空判断不可少
	for cursor != nil && cursor.data != data {
		cursor = cursor.next
		index++
	}

	if cursor == nil {
		return -1
	} else {
		return index
	}
}

//按值删除
func (l *linkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	previousToDelete = nil
	l.length--
}

//删除结点
func (l *linkedList) DeleteNode(p *node) {
	if l.length == 0 || p == nil {
		return
	}

	previousToDelete := l.head
	for previousToDelete.next != p {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	p = nil
	l.length--
}

func (l *linkedList) InsertBefore(n *node) {
	if l.head == nil || n == nil {
		return
	}

	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//在某个结点后面插入结点
func (l *linkedList) InsertAfter(p *node, value int) bool {
	if p == nil {
		return false
	}
	//1.生成一个新结点且数据域赋值
	node := initNode(value)
	old := p.next
	//新结点插入到尾结点之后
	p.next = node
	node.next = old
	l.length++
	return true
}
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ->", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
func main() {
	mylist := initLinkedList()
	mylist.InsertBefore(initNode(1))
	mylist.InsertBefore(initNode(2))
	mylist.InsertBefore(initNode(3))
	mylist.InsertBefore(initNode(4))
	mylist.InsertBefore(initNode(5))
	mylist.InsertBefore(initNode(6))
	mylist.InsertBefore(initNode(7))
	mylist.printListData()

	/*
		ele := mylist.GetElemWithIndex(1)
		if ele != nil {
			fmt.Println(ele.data)
		}

		ele2 := mylist.FindWithData(9)
		if ele2 != nil {
			fmt.Println(ele2.data)
		}

	index := mylist.FindData(7)
	fmt.Println(index)


	*/
	index2 := mylist.LocateElem(90)
	fmt.Println(index2)
	index3 := mylist.LocateElem(70)
	fmt.Println(index3)
}
