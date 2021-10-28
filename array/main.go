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

	//数组的大小
	theSize int
}

func NewArray() *Array  {
	//初始化结构体
	array := new(Array)

	//开辟内在空间
	array.dataStore = make([]interface{},0,10)

	array.theSize = 0

	return array
}

func (array *Array) Size()int  {
	return array.theSize
}

func (array *Array) Get(index int)(interface{},error) {
    if index < 0 || index >= array.theSize {
		return nil,errors.New("Index out of bounds")
	}
	return array.dataStore[index],nil
}

func (array *Array) Append(newval interface{}) {
	//叠加数据
	array.dataStore = append(array.dataStore,newval)
	array.theSize++
}
func (array *Array) String() string {
	return fmt.Sprint(array.dataStore)
}
func main()  {
	array := NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5)
	array.Append(6)
	array.Append(7)
	fmt.Println(array.String())


}
