package hascycle

type ListNode struct {
	Val int
	Next *ListNode
}
 
func HasCycle(head *ListNode) bool {
	used := make(map[*ListNode]*ListNode)
    for head != nil && head.Next != nil {
        node := used[head]
        if node == head {
            return true
        }
        used[head] = head
        head = head.Next
    }
    return false
}

func DetectCycle(head *ListNode) *ListNode {
    used := make(map[*ListNode]*ListNode)
    for head != nil && head.Next != nil {
        node := used[head]
        if node == head {
            return node
        }
        used[head] = head
        head = head.Next
    }
    return nil
}