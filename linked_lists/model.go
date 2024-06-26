package linked_lists

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
