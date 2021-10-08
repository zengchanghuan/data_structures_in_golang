package main

import "fmt"

func a() func()int  {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func testC1()  {
	c := a()
	fmt.Println(c())
}
func test()func()  {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x,x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
func testC2()  {
	f := test()
	f()
}
func add(base int) func(int) int  {
	return func(i int) int {
		base += 1
		return base
	}
}
func testC3()  {
	temp1 := add(10)
	fmt.Println(temp1(1),temp1(2))

	temp2 := add(100)
	fmt.Println(temp2(1),temp2(2))
}

type Test struct {
	name string
}

func (t *Test)Close()  {
	fmt.Println(t.name,"closed")
}

func Close(t Test)  {
	t.Close()
}

func test2(x int)  {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")

}
func main()  {
	test2(5)
}
