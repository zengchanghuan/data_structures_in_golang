package main

import "fmt"

func MergeSort(list []int)[]int  {
	if len(list) < 2 {
		return list
	}

	mid := len(list) / 2
	left := list[:mid]
	right := list[mid:]

	return merge(MergeSort(left),MergeSort(right))

}
func merge(l []int, r []int) []int {
	var result []int

	for len(l) > 0 || len(r) > 0 {
		if len(l) > 0 && len(r) > 0 {
			if l[0] < r[0] {
				result = append(result, l[0])
				l = l[1:]
			} else {
				result = append(result, r[0])
				r = r[1:]
			}
		} else if len(l) > 0 {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}
func main()  {
	list := []int{91,28,73,46,55,64,37,82,19,-90}
	fmt.Println("before:", list)

	fmt.Println("after: ", MergeSort(list))
}
