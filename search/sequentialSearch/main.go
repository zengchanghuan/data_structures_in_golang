package main

import "fmt"
//如果线性表无序，采用顺序查找；线性表有序，采用折半查找
func seqSearch(key int,a []int)int  {
	if a[0] == key {
		return 0
	}
	a[0] = key
	i := len(a) - 1
	for a[i] != key {
		i --
	}
	return i
}

func main()  {
	a := []int{0,-1,5,3,232,9,32,535}
	fmt.Println(seqSearch(32,a))
}
