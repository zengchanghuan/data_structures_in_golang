package main

import "fmt"

type list struct {
	data string
	//代表每个结点能够访问的结点
	next []*list
}

var m map [string]bool

func DFS(row []*list)  {
	if len(row) == 0 {
		return
	}

	for i := range row {
		if ok,_ := m[row[i].data]; !ok {
			fmt.Print(row[i].data," ")
			m[row[i].data] = true
			if len(row[i].next) > 0 {
				DFS(row[i].next)
			}
		}
	}
}

func BFS(row []*list)  {
	if len(row) == 0 {
		return
	}

	newRow := make([]*list,0)

	for i := range row {
		if ok,_ := m[row[i].data];!ok {
			fmt.Print(row[i].data," ")
			m[row[i].data] = true
			if len(row[i].next) != 0 {
				newRow = append(newRow,row[i].next...)
			}
		}
	}
	BFS(newRow)
}

func main()  {
	m = make(map[string]bool)
	v0 := &list{data: "v0"}
	v1 := &list{data: "v1"}
	v2 := &list{data: "v2"}
	v3 := &list{data: "v3"}
	v4 := &list{data: "v4"}
	v5 := &list{data: "v5"}
	v6 := &list{data: "v6"}

	v0.next = append(v0.next,v1,v2,v3)
	v2.next = append(v2.next,v4)
	v1.next = append(v1.next,v4,v5)
	v3.next = append(v3.next,v5)
	v4.next = append(v4.next,v6)
	v5.next = append(v5.next,v6)
	fmt.Print("DFS : ")
	DFS([]*list{v0})

	m = make(map[string]bool)
	fmt.Print("\nBFS : ")
	BFS([]*list{v0})

}