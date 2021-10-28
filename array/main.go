package main

import (
	"errors"
	"fmt"
)

type ArrayList interface {
	//数组大小
	Size() int

	//第几个元素
	Get(index int) (interface{},error)

	//修改数据
	Set(index int,newval interface{})error

	//插入数据
	Insert(index int,newval interface{})error

	//追加
	Append(newval interface{})

	//清空数据
	Clear()

	//删除
	Delete(index int)error

	//返回字符串
	String()string

}

type Array struct {
	//数组存储
	dataStore [] interface{}

	//数组的大小 首字母小写私有，只能内部用，大写公有可供外部用
	TheSize int
}

func NewArray() *Array  {
	//初始化结构体
	array := new(Array)

	//开辟内在空间
	array.dataStore = make([]interface{},0,10)

	array.TheSize = 0

	return array
}

func (array *Array)checkIsFull()  {
	if array.TheSize == cap(array.dataStore) {
		//开辟双倍内存
		newDataStore := make([]interface{},2 * array.TheSize,2 * array.TheSize)
		//拷贝
		copy(newDataStore,array.dataStore)
		//for i := 0; i < len(array.dataStore); i++ {
		//	newDataStore[i] = array.dataStore[i]
		//}
		//赋值
		array.dataStore = newDataStore
		fmt.Println(array.dataStore)
		fmt.Println(newDataStore)
	}
}

func (array *Array) Size()int  {
	return array.TheSize
}

func (array *Array) Get(index int)(interface{},error) {
    if index < 0 || index >= array.TheSize {
		return nil,errors.New("Index out of bounds")
	}
	return array.dataStore[index],nil
}

func (array *Array) Set(index int,newval interface{})error {
	if index < 0 || index >= array.TheSize {
		return errors.New("Index out of bounds")
	}

	//设置数据
    array.dataStore[index] = newval
	return nil

}
func (array *Array) Append(newval interface{}) {
	//叠加数据
	array.dataStore = append(array.dataStore,newval)
	array.TheSize++
}
func (array *Array) String() string {
	return fmt.Sprint(array.dataStore)
}

func (array *Array) Clear() {
	array.dataStore = make([]interface{},0,10)
	array.TheSize = 0
}

//重点
func (array *Array) Insert(index int,newval interface{})error {
	if index < 0 || index >= array.TheSize {
		return errors.New("Index out of bounds")
	}
	//检测内存,如果满了会自动追加
	array.checkIsFull()

	//插入数据，内存移动一位
	array.dataStore = array.dataStore[:array.TheSize+1]

	for i := array.TheSize;i > index; i-- {
		array.dataStore[i] = array.dataStore[i - 1]
	}
	array.dataStore[index] = newval
	array.TheSize++

	return nil
}

//重点
func (array *Array) Delete(index int)error  {
	//重新叠加，跳过删除
	array.dataStore = append(array.dataStore[:index],array.dataStore[index + 1:]...)
	array.TheSize--
	return nil
}
func main()  {
	array := NewArray()
	array.Append("a")
	array.Append("b")
	array.Append("c")
    for i := 0; i < 15; i++ {
		array.Insert(1,"x")
	}
	fmt.Println(array)
}
