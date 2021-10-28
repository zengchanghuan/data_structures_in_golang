package main

import "fmt"

type node struct {
	value int
	next *node
}

type circle struct {
	tail *node
	length int
}
func (c *circle) add(value int) {
	newNode := &node{value, nil}
	if c.length == 0 { //空链表中添加节点
		c.tail = newNode
		c.tail.next = newNode
	} else {
		newNode.next = c.tail.next
		c.tail.next = newNode
		c.tail = newNode
	}
	c.length += 1
	//c.printCircle()
}

// 删除节点：
func (c *circle) remove(v int) {
	if c.length == 0 {
		fmt.Println("空环")
		return
	} else if c.length == 1 && c.tail.value == v { //链表中只有一个节点的特殊情况
		c.tail = nil
		c.length = 0
		//c.printCircle()
		return
	}
	pre := c.tail
	cur := c.tail.next // 头节点
	for i := 0; i < c.length; i++ {
		if cur.value == v {
			if cur == c.tail { //如果删除的节点是尾节点,需更新tail
				c.tail = pre
			}
			pre.next = cur.next
			c.length -= 1
			//c.printCircle()
			return
		}
		pre = cur
		cur = cur.next
	}
	fmt.Println(v, "不在环中")
}
//打印节点：
func (c *circle) printCircle() {
	if c.length == 0 {
		fmt.Println("空环")
		return
	}
	cur := c.tail.next // 头节点
	for i := 0; i < c.length; i++ {
		fmt.Printf("%d ", cur.value)
		cur = cur.next
	}
	fmt.Println()
}

func main()  {
	circle := new(circle)
	circle.add(1)
	circle.add(2)
	circle.add(3)
	circle.add(4)
	circle.add(5)
	circle.add(6)
	circle.add(7)
	circle.printCircle()
	circle.remove(32)
	circle.printCircle()
}