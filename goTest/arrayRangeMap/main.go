package main

import "fmt"

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
func testArray()  {
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println(numbers[:3])

	fmt.Println(numbers[4:])
}

func appendAndCopyTest()  {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers,5)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func rangeTest()  {
	nums := []int{2,3,4}
	sum := 0

	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，
	//所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i,num := range nums {
		if num == 3 {
			fmt.Println("index",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banna"}
	for k,v := range kvs {
		fmt.Printf("%s %s\n",k,v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i,c)
	}
}

func mapTest()  {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap ["France"] = "巴黎"
	countryCapitalMap ["Italy"] = "罗马"
	countryCapitalMap ["Japan"] = "东京"
	countryCapitalMap ["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital,ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func mapDeleteTest()  {
	countryCapitalMap := map[string]string{"France":"Paris","Italy":"Rome","" +
		"Japan":"Tokyo","India":"New delhi"}

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
	delete(countryCapitalMap,"France")
	fmt.Println("删除元素后地图")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}
func main()  {
	testArray()
	appendAndCopyTest()
	rangeTest()
	mapTest()
	mapDeleteTest()
}
