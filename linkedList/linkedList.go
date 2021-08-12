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

//取单链表中第i个元素
func (l *linkedList) GetElem(index int) *node {
	index -= 1
	if l.length == 0 || index > l.length {
		return nil
	}

	//从链表头指值出发，顺着链域next逐个往下搜索
	cur := l.head
	var i = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//查找，按值查找，根据数据获取该数据所在的地址
func (l *linkedList) FindByData(data int) *node {
	if l.length == 0 {
		return nil
	}

	cur := l.head.next
	for {
		if data != cur.data {
			cur = cur.next
		}
		return cur
	}
	return nil
}

//查找，按值查找，根据数据获取该数据所在的位置序号
func (l *linkedList) FindData(data int) int {
	if l.length == 0 {
		return -1
	}

	cur := l.head.next
	index := 1
	for {
		if data != cur.data {
			cur = cur.next
			index++
		}
	}
	if cur != nil {
		return index
	} else {
		return -1
	}
}

func (l *linkedList) withinRange(index int) bool {
	return index >= 0 && index < l.length
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

//在某一结点前插入一个结点
//在某一结点前插入一个结点
func (l *linkedList) InsertBeforeC(p *node) {
	if l.head == nil || p == nil {
		return
	}

	//首元结点
	cur := l.head.next
	//头结点
	pre := l.head

	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return
	}

	pre.next = p
	p.next = cur
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

	ele := mylist.GetElem(2)
	fmt.Println(ele.data)
}
