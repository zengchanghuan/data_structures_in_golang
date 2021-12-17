package main

import "fmt"

func main()  {
	func(a,b int) {
		sum := a + b
		fmt.Println(sum)
	}(10,30)

	f := func(a,b int) {
		sum := a + b
		fmt.Println(sum)
	}

	f(10,20)
}
