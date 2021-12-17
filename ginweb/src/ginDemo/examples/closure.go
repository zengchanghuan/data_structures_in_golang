package main

import "fmt"

func getSequence() func() int  {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main()  {
	f1 := getSequence()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	f2 := getSequence()
	fmt.Println(f2())
	fmt.Println(f2())
}