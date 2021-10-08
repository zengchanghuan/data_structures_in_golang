package main

import (
	"fmt"
	"sync"
)
//并发执行 随机调试
var wg sync.WaitGroup
func hello(i int)  {
	defer wg.Done()
	fmt.Println("hello goroutine!",i)
}
func main()  {
	for i := 0;i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
