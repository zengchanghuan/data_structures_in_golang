package main

import "fmt"

func valueIntTest(a int) int  {
	return a + 10
}

func pointerIntTest(a *int) int  {
	return *a + 10
}

func structTestValue()  {
	a := 2
	fmt.Println(valueIntTest(a))

	b := 5
	fmt.Println(&b)
	fmt.Println(pointerIntTest(&b))

}

type PersonD struct {
	id int
	name string
}

func (p PersonD) valueShowName()  {
	fmt.Println(p.name)
}

func (p *PersonD) pointShowName()  {
	fmt.Println(p.name)
}

func structTestFunc()  {
	//personValue := PersonD{101,"hello world"}
	//personValue.valueShowName()
	//personValue.pointShowName()

	personPointer := &PersonD{102,"hello golang"}
	personPointer.valueShowName()
	personPointer.pointShowName()

}
func main()  {
	structTestFunc()
}
