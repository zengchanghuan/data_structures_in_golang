package main

import "fmt"
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move()  {
	fmt.Println("猫会动")
}

func main() {
	var x interface{}
	x = "pprof.cn"
	v,ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型判断失败")
	}
}
