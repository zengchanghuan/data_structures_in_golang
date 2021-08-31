package main

import (
	"fmt"
)

func QuickSort(a []int)  {
	if len(a) < 2 {
		return
	}
	quicksort(a,0,len(a) - 1)
}

func quicksort(values []int, left int,right int)  {
	if left >= right {
		return
	}

	pivot := values[left]
	i := left + 1

	for j := left; j <= right; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[left], values[i-1] = values[i-1], pivot

	quicksort(values, left, i-2)
	quicksort(values, i, right)

}

func main()  {
	var list = []int{91,28,73,46,55,64,37,82,19}
	fmt.Println("unsorted:", list)
	QuickSort(list)
	fmt.Println("sorted:  ", list)
}
