package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// zero or one node in the list
		return head
	}

	prev, next := head, head.Next
	for next != nil {
		tmp := next.Next
		next.Next = prev
		prev = next
		next = tmp
	}
	head.Next = nil
	head = prev

	return head
}
