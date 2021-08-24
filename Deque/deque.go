package main

import "fmt"

type MyCircularDeque struct {
	queue []int
	head int
	tail int
}

func InitDeque(k int) MyCircularDeque  {
	return MyCircularDeque{
		queue: make([]int,k + 1),
		head: 0,
		tail: 0,
	}
}

/*
头进 头部进队，先将this.head前移一位，然后将元素先放入this.head位置
队头前移一位即this.head - 1
 */
func (this *MyCircularDeque)InsertFront(value int) bool  {
	if this.IsFull() {
		return false
	}

	this.head = (this.head - 1 + len(this.queue)) % len(this.queue)
	this.queue[this.head] = value
	return true
}

//尾进：尾部进队，即后端进队，先将元素放入队尾位置，然后队尾后移一位，后移进为处理边界情况，需要加1后模len(this.queue)取余
func (this *MyCircularDeque) InsertLast(value int)bool  {
	if this.IsFull() {
		return false
	}

	//先放入尾部
	this.queue[this.tail] = value
	this.tail = (this.tail + 1) % len(this.queue) //向后移动一位
	return true
}
/*
头出 即前端出队
 */
func (this *MyCircularDeque) DeleteFront()bool  {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.queue)//向后移动一位

	return true
}

/*尾出 尾部出队，即后端出队，先将this.tail前移一位，然后取出元素，前移一位即this.tail - 1，
当this.tail为零时，this.tail - 1为负值，因此加上len(this.queue)，正好是len(this.queue)-1的位置。
this.tail - 1为正值时，加上len(this.queue)就超过了下标范围，需要模len(this.queue)取余
 */
func (this *MyCircularDeque) DeleteLast()bool  {
	if this.IsEmpty() {
		return false
	}

	//this.tail = (this.tail - len(this.queue)) % len(this.queue)
	this.tail = (this.tail - 1 + len(this.queue)) % len(this.queue)

	return true
}

func (this *MyCircularDeque) GetFront()int  {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

//取队尾时，尾指针不移动
func (this *MyCircularDeque) GetRear()int  {
	if this.IsEmpty() {
		return -1
	}
	n := (this.tail - 1 + len(this.queue)) % len(this.queue)
	return this.queue[n]
}
func (this *MyCircularDeque)IsEmpty()bool  {
	return this.head == this.tail
}
func (this *MyCircularDeque) IsFull() bool  {
	//队尾后移一位等于队头，表明队满
	return (this.tail + 1) % len(this.queue) == this.head
}

func (this *MyCircularDeque) Length()int  {
	return len(this.queue)
}

func (this *MyCircularDeque) Traverse()  {
	if this.IsEmpty() {
		fmt.Println("deque is empty")
		return
	}
	temp := this.GetFront()
	for temp != this.tail {
		fmt.Printf("%d\n",this.queue[temp])
		temp = (temp+1) % len(this.queue)
	}
	fmt.Println("Traverse is over")
}
func main()  {
	deque := InitDeque(6)
	deque.InsertFront(99)
	deque.InsertFront(2)
	deque.InsertFront(3)
	deque.InsertFront(4)
	deque.InsertFront(5)
	deque.InsertFront(234)
	fmt.Println(deque)
	fmt.Println(deque.GetFront())
	fmt.Println(deque.GetRear())
	fmt.Println(deque.DeleteLast())
	fmt.Println(deque.GetFront())
	fmt.Println(deque.GetRear())
	fmt.Println(deque)
}
