package linked_lists

func (ln *ListNode) reverseListNode() *ListNode {
	var prev *ListNode
	curr := ln
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}