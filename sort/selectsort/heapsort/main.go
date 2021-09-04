package main

import (
	"fmt"
)
// 堆排
func heapSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	maxHeap(nums, len(nums) - 1)

	for i := len(nums) - 1; i >= 0; i -- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums[:i], 0, len(nums[:i]) - 1)
	}

	return nums
}


func maxHeap(nums []int, length int){
	for i := len(nums)>> 1 - 1; i >= 0; i -- {
		heapify(nums, i, length)
	}
}

func heapify(nums []int, i int, len int){
	left := 2 * i + 1
	right := 2 * i + 2
	large := i

	if left <= len && nums[left] > nums[large] {
		large = left
	}

	if right <= len && nums[right] > nums[large] {
		large = right
	}

	if large != i {
		nums[large], nums[i] = nums[i], nums[large]
		heapify(nums, large, len)
	}
}

func main()  {
	list := []int{91,28,73,46,55,64,37,82,19,-90}
	fmt.Println("before:", list)

	fmt.Println("after: ", heapSort(list))
}
