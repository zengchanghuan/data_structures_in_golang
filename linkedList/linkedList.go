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
	//if l.length > 0 {
	//	return true
	//} else {
	//	return false
	//}
	if l.head.next == nil {
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

	cursor := l.head
	for {
		if cursor.data == data {
			break
		}
		if cursor.next == nil {
			break
		}
		cursor = cursor.next
	}
	return cursor
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

//在第i个结点前插入一个结点
func (l *linkedList) InsertBefore(index, value int) {
	if l.head == nil || index < 1 {
		return
	}

	cursor := l.head
	j := 0
	//寻找第i-1个结点，cursor指各i-1个结点
	for cursor != nil && j < index-1 {
		j++
		cursor = cursor.next
	}
	if cursor == nil {
		return
	} else {
		s := initNode(value)
		s.next = cursor.next
		cursor.next = s
	}
}

//删除第i个元素
func (l *linkedList) DeleteIndex(i int) {
	if i < 1 {
		return
	}
	pre := l.head
	j := 0
	for j < i-1 && pre != nil {
		j++
		pre = pre.next
	}
	if pre == nil {
		return
	} else {
		if pre.next == nil {
			return
		}
		pre.next = pre.next.next
		return
	}
}

//按值删除
func (l *linkedList) DeleteWithValue(value int) {
	if l.head == nil || l.length == 0 {
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

func (l *linkedList) InsertPre(n *node) {
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
	mylist.InsertPre(initNode(1))
	mylist.InsertPre(initNode(2))
	mylist.InsertPre(initNode(3))
	mylist.InsertPre(initNode(4))
	mylist.InsertPre(initNode(5))
	mylist.InsertPre(initNode(6))
	mylist.InsertPre(initNode(7))
	mylist.printListData()

	//mylist.InsertBefore(3,54)
	//mylist.DeleteIndex(1)
	//mylist.DeleteWithValue(3)
	mylist.printListData()
	//mylist.DeleteWithValue(7)
	mylist.printListData()

	d := mylist.GetElemWithIndex(3)
	mylist.DeleteNode(d)
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



		index2 := mylist.LocateElem(90)
		fmt.Println(index2)
		index3 := mylist.LocateElem(70)
		fmt.Println(index3)
	*/

}
