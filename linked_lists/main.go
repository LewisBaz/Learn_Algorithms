package linked_lists

func Run() {
	checkAddTwoNumbers()
}

func checkAddTwoNumbers() {
	head1 := &ListNode{Val: 2}
	head1.Next = &ListNode{Val: 4}
	head1.Next.Next = &ListNode{Val: 3}

	head2 := &ListNode{Val: 5}
	head2.Next = &ListNode{Val: 6}
	head2.Next.Next = &ListNode{Val: 4}

    printList(addTwoNumbers(head1, head2))
}