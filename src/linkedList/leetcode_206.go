package linkedList

/**
链表反转
*/
type ListNode struct {
	Next *ListNode
	Val  int
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
