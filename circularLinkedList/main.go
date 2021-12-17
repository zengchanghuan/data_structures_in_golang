package main

import (
	"container/ring"
	"fmt"
)

func main()  {
	r := ring.New(5)
	for i := 0;i < 5; i++ {
		r.Value = i
		r = r.Next()//循环下一个
	}
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	r3 := r.Move(5)
	fmt.Println(r3)//第三个元素的地址
	fmt.Println(r)//第一个元素的地址
	r3.Do(func(i interface{}) {
		fmt.Println(i)
	})

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
	r1 := r.Link(r3)
	fmt.Println(r1)
	fmt.Println(r)

}
