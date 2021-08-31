package main

import "fmt"

func bubbleSort(input []int)  {
	len := len(input)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
}
func main()  {
	input := []int{100000, 5, 7, 9, 44, 33, 4, 6, 882222226, 453, 244, 234}

	fmt.Println(input)

	bubbleSort(input)

	fmt.Println(input)
}
