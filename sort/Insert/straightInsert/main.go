package main

func InsertSort2(arr []int,l int)  {
	for i := 1; i < l; i++ {
		vi := arr[i]
		j := i
		for ;j > 0 && arr[j-1] >= vi; j-- {
			arr[j] = arr[j - 1]
		}
		arr[j] = vi
	}
}
/*
在插入排序中，我们从前到后依次处理未排序好元素，
对于每个元素，我们将它与之前排好序的元素进行比较，
找到对应的位置后并插入

基本思想：

步骤：
1.从第二个元素开始，从后向前扫描之前的元素序列
2.如果当前扫描的元素大于新元素，将扫描元素移动到下一位
3.重复步骤2，直到找到一个小于或等于新元素的位置
4.将新元素插入到该位置
5.对于之后的元素重复步骤1-4
基于此有折半查找
 */
func InsertSort(arr []int)  {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		//从后往前扫开始的位置
		j := i -1
		//被扫描元素大于新元素，就进行移位
		for ;j >= 0 && arr[j] > cur; j-- {
			arr[j+1] = arr[j]
		}
		//将新元素插入该位置
		arr[j+1] = cur
	}
}
func main()  {
	arr := []int{49, 38, 65, 97, 76, 13, 27, 49}
	//fmt.Println(arr)
	InsertSort(arr)
	//fmt.Println(arr)
}
