package main

import "fmt"
/*
先将数组分割成若干个子序列
分别进行插入排序，待整个序列是的记录基本有序时，
再对全体记录进行一次直接插入排序
特点：
1）缩小增量
2）多遍插入排序
3）增量序列必须是递减的
4）增量序列应该是互质的 15 7 3 1 inc = 2^k-1
 */
func shellsort(arr []int) {
	//inc相当于是步长因子
	for inc := len(arr) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
		for i := inc; i < len(arr); i++ {
			j, temp := i, arr[i]
			for ; j >= inc && arr[j-inc] > temp; j -= inc {
				fmt.Printf("inc = %d\n",inc)
				arr[j] = arr[j-inc]
			}
			arr[j] = temp
		}
	}
}
//缩小增量插入排序
func main()  {

	arr := []int{90,-80,70,-60,50,-40,30,-20,10,0}
	fmt.Println("before:", arr)
	shellsort(arr)
	fmt.Println("after: ", arr)
}