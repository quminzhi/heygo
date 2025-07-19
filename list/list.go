package list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206
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

// 141
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

// 83
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

// 203
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

// 876
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

// 92
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

// 143
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Reverse the right half
	// e.g., o->o<->o
	mid, end := head, head
	for end != nil && end.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}

	cur, next := mid, mid.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		// Update working pointers
		cur = next
		next = tmp
	}

	// Reorder the list
	begin, end := head, cur
	dummy := &ListNode{}
	last := dummy
	for {
		if begin != mid {
			// Append begin
			last.Next = begin
			last = last.Next
			begin = begin.Next
		}
		if end != mid {
			// Append end
			last.Next = end
			last = last.Next
			end = end.Next
		}
		if begin == mid && end == mid {
			// Break on both hit the mid
			break
		}
	}
	// Append the mid
	last.Next = mid
	last = last.Next
	last.Next = nil
	head = dummy.Next
}

// 82
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	last := dummy
	cur := head // Points to the node to validate
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		if cur.Next == next {
			// A distinct number found
			last.Next = cur
			last = last.Next
		}
		cur = next
	}
	last.Next = nil
	return dummy.Next
}

// 86
func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	lessLast := less
	more := &ListNode{}
	moreLast := more
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			lessLast.Next = cur
			lessLast = lessLast.Next
		} else {
			moreLast.Next = cur
			moreLast = moreLast.Next
		}
	}
	lessLast.Next = more.Next
	moreLast.Next = nil
	return less.Next
}

// 61
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	shift := k % n
	if shift == 0 {
		return head
	}

	// Find the last node in the rotated list
	last := head
	for i := 0; i < n-shift-1; i++ {
		last = last.Next
	}
	newHead := last.Next
	last.Next = nil

	// Find the last in the original list and connect
	for last = newHead; last.Next != nil; last = last.Next {
	}
	last.Next = head
	head = newHead
	return head
}

// 147
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt32, Next: nil}

	cur := head
	for cur != nil {
		next := cur.Next

		// Find the right place for each node to insert
		ins := dummy
		for ; ins.Next != nil && ins.Next.Val < cur.Val; ins = ins.Next {
		}
		// Insert after ins
		tmp := ins.Next
		ins.Next = cur
		cur.Next = tmp

		cur = next
	}
	return dummy.Next
}

type Node struct {
	Value  int
	Next   *Node
	Random *Node
}

// 138
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	// Clone points
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Value: cur.Value, Next: nil}
	}

	// Clone pointers
	for k, v := range m {
		v.Next = m[k.Next]
		v.Random = m[k.Random]
	}

	return m[head]
}

// 24
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	p, q := head, head.Next
	for p != nil && q != nil {
		next := q.Next

		// Swap p and q
		prev.Next = q
		p.Next = next
		q.Next = p
	
		// Update prev, p and q
		prev = p
		p = next
		if next == nil {
			break
		}
		q = next.Next
	}

	return dummy.Next
}
