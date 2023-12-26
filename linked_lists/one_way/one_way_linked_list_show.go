package oneway

import "fmt"

var top = &OneWayLinkedCell{Value: 44}

func Show() {
	top.AddInEnd(56)
	top.AddInEnd(77)
	top.AddInEnd(33)
	top.AddInEnd(54)
	top.AddAfter(77, 99)
	top.AddAfter(34, 35)
	top = top.Remove(44)
	top.ShowAll()
	fmt.Println(top.FindNextAfter(54))
}