package main

import "fmt"

/*
为充分利用向量空间，克服"假溢出"现象的方法是：
将向量空间想象为一个首尾相接的圆环，并称这种向量为循环向量。
存储在其中的队列称为循环队列（Circular Queue）。
循环队列是把顺序队列首尾相连，把存储队列元素的表从逻辑上看成一个环，成为循环队列。

*/
type MyCircularQueue struct {
	queue []int
	head int
	count int //队列实际大小
	capacity int //队列总长度
}

func InitCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int,k,k),
		head: 0,
		count: 0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count >= this.capacity {
		return false
	}
	//取模获取入队下标
	this.queue[(this.head + this.count) % this.capacity] = value
	this.count++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}

	//删除数组的头
	this.head = this.head + 1 % this.capacity

	this.count--
	return true
}

// 从队首获取元素。如果队列为空，返回 -1 。
func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.head]
}

//获取队尾元素。如果队列为空，返回 -1 。
func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}

	return this.queue[this.head + this.count - 1] % this.capacity
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.capacity
}

func main()  {
	cq := InitCircularQueue(6)
	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.EnQueue(3)
	cq.EnQueue(4)
	cq.EnQueue(5)
	cq.EnQueue(77)

	//fmt.Println(cq.IsFull())
	fmt.Println(cq.Front())
	fmt.Println(cq.Rear())
	//cq.EnQueue(9)
	//
	//
	//fmt.Println(cq.DeQueue())
	//fmt.Println(cq.DeQueue())
}