package main

import "fmt"

func BinarySearch(key int,arr []int)int  {
	low,high := 0,len(arr) - 1
	for low <= high{
		mid := (low + high) >> 1
		if arr[mid] > key {
			high = mid - 1
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func SearchRecursion(key int,arr []int,low,high int)int  {
	if low > high {
		return -1
	}

	mid := (low + high) >> 1
	if arr[mid] > key {
		return SearchRecursion(key,arr,low,mid - 1)
	} else if arr[mid] < key {
		return SearchRecursion(key,arr,mid + 1,high)
	} else {
		return mid
	}
}
func main()  {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	tests := []int{1, 7, 5, 10, -29}
	for _,v := range tests {
		//fmt.Println("BinarySearch",BinarySearch(v,a))

		fmt.Println("SearchRecursion",SearchRecursion(v,a,0,len(a) - 1))

	}


}
