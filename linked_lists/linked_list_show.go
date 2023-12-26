package linked_lists

import (
	"fmt"
	ll "main/linked_lists/cell"
)

func ShowLinkedList() (string) {
	ll := &ll.LinkedCell{}
	ll.Add(11) 
	fmt.Println(ll.Next.Value)
	ll.Add(17)
	fmt.Println(ll.Next.Value)
	ll.Add(28)
	fmt.Println(ll.Next.Value)
	fmt.Println(ll.Show())
	return ll.Show()
}