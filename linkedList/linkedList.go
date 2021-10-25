package main

import "fmt"

/*
单链表的基本操作：初始化，创建，取值，查找，插入，删除
*/
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
	if l.head.next == nil {
		return true
	} else {
		return false
	}
}

func (l *linkedList) Last() *node  {
	h := l.head
	for h != nil && h.next != nil {
		h = h.next
	}
	return h
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
		l.head = &node{0,nil}
	}
	l.length = 0
	l.head.next = nil

}

//表长
func (l *linkedList) Length() int {
	if l.head == nil {
		return 0
	}
	return l.length
}

//查找，按值查找返回位置序号
func (l *linkedList) LocateElem(data int) int {
	if l.head == nil || l.length == 0 {
		return -1
	}
	cursor := l.head
	index := 1
	//如果没有找到，cursor.next 为空，故下面cursor为空判断不可少
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
//查找，按位查找 返回地址
func (l *linkedList) GetElemWithIndex(index int) *node {
	//要不要都行，从0开始取可以不要
	//index -= 1
	if index < 0 {
		fmt.Println("index is error")
		return nil
	}
	if l.length == 0 || index > l.length {
		return nil
	}

	//从链表头指值出发，顺着链域next逐个往下搜索
	cursor := l.head
	for i:= 0; i < index; i++ {
		cursor = cursor.next
	}
	return cursor
}

//在第i个结点前插入一个结点
func (l *linkedList) InsertBefore(index, value int) {
	if l.head == nil || index < 1 {
		return
	}

	L := l.head
	j := 0
	//寻找第i-1个结点，cursor指向i-1个结点
	for L != nil && j < index-1 {
		j++
		L = L.next
	}
	if L == nil {
		return
	} else {
		s := initNode(value)
		//以下两行是关键代码 新结点的下一个结点存储是头结点，实现头插法的关键代码
		s.next = L.next
		//设置新的头结点，实现下一步头插法
		L.next = s
	}
}
//头插法 逆序建表
func (l *linkedList) InsertPre(n *node) {
	if l.head == nil || n == nil {
		return
	}

	//以下两种写法二选一，第一种好理解，第二种更简洁
	/*
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
	*/
	n.next = l.head
	l.head = n
	l.length++

}

//尾插法 正序建表
func (l *linkedList) InsertAfter(node *node) {
	if node == nil {
		return
	}
	if l.length == 0 {
		l.head = node
	} else {
		//取出头部结点
		L := l.head
		//判断是否有下一个结点
		for L.next != nil {
			L = L.next
		}
		L.next = node
	}
	l.length++
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
	fmt.Println(mylist.LocateElem(523))


	mylist2 := initLinkedList()
	mylist2.InsertAfter(initNode(1))
	mylist2.InsertAfter(initNode(2))
	mylist2.InsertAfter(initNode(3))
	mylist2.InsertAfter(initNode(4))
	mylist2.InsertAfter(initNode(5))
	mylist2.InsertAfter(initNode(6))
	mylist2.InsertAfter(initNode(7))
	mylist2.printListData()

	//fmt.Println(mylist.GetElemWithIndex(4))
	fmt.Println(mylist2.LocateElem(523))


	//mylist.InsertBefore(3,54)
	//mylist.DeleteIndex(1)
	//mylist.DeleteWithValue(3)
	//mylist.printListData()
	//mylist.DeleteWithValue(7)
	//mylist.printListData()
	/*
	d := mylist.GetElemWithIndex(3)
	mylist.DeleteNode(d)
	mylist.printListData()


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
