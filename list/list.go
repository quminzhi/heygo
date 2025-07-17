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

	cur, next := head, head.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	head.Next = nil
	head = cur

	return head
}

// hasCycle detects if a linked list has a circle in it
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}

	return false
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for ; cur != nil && cur.Val == val; cur = cur.Next {
	}

	head, prev := cur, cur
	for ; cur != nil; cur = cur.Next {
		if cur.Val != val {
			prev = cur
		} else {
			prev.Next = cur.Next
		}
	}
	return head
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid, last := head, head
	for mid != nil && last != nil && last.Next != nil {
		mid = mid.Next
		last = last.Next.Next
	}
	return mid
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// Move to the left
	cur := dummy
	prev := dummy // the one before left
	for i := 0; i < left; i++ {
		prev = cur
		cur = cur.Next
	}

	// Reverse the nodes between left and right
	next := cur.Next
	for i := 0; i < right-left; i++ {
		tmp := next.Next
		next.Next = cur
		// Update working pointers
		cur = next
		next = tmp
	}

	// Connect
	prev.Next.Next = next
	prev.Next = cur

	return dummy.Next
}
