package linked_lists

import "fmt"

type LinkedCell struct {
	Value int
	Next *LinkedCell
}

func (cell *LinkedCell) Add(integer int) {
	cell.Next = &LinkedCell{Value: integer, Next: nil}
}

func (cell *LinkedCell) Show() (string) {
	res := ""
	for cell.Next != nil {
		res += string(rune(cell.Value))
	}
	fmt.Println(res)
	return res
}