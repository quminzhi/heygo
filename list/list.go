package list

import (
	"container/heap"
	"math"
)

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

// 328
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyOdd := &ListNode{}
	lastOdd := dummyOdd
	dummyEven := &ListNode{}
	lastEven := dummyEven
	for cur, i := head, 1; cur != nil; i = i + 1 {
		next := cur.Next
		if i%2 == 0 {
			// Even
			tmp := lastEven.Next
			lastEven.Next = cur
			cur.Next = tmp
			lastEven = cur
		} else {
			// Odd
			tmp := lastOdd.Next
			lastOdd.Next = cur
			cur.Next = tmp
			lastOdd = cur
		}
		cur = next
	}
	lastOdd.Next = dummyEven.Next
	return dummyOdd.Next
}

// 725
// no two parts should have a size differing by more than one
func splitListToParts(head *ListNode, k int) []*ListNode {
	splitList := make([]*ListNode, k)
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	cur := head
	for group := 0; group < k; group++ {
		groupSize := n / k
		if group+1 <= n%k {
			groupSize++
		}
		splitList[group] = cur
		// group number may be greater than the number of the list
		if cur == nil {
			continue
		}
		// Guarantee that cur always points to some node in the list
		for i := 0; i < groupSize-1; i++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}

	return splitList
}

// 21
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	last := dummy
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			next := l1.Next // Backup
			last.Next = l1
			l1.Next = nil
			last = l1
			l1 = next
		} else {
			next := l2.Next
			last.Next = l2
			l2.Next = nil
			last = l2
			l2 = next
		}
	}
	for l1 != nil {
		last.Next = l1
		last = l1
		l1 = l1.Next
	}
	for l2 != nil {
		last.Next = l2
		last = l2
		l2 = l2.Next
	}
	last.Next = nil
	return dummy.Next
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l >= r {
		return lists[l]
	}
	mid := l + (r-l)/2
	left := merge(lists, l, mid)
	right := merge(lists, mid+1, r)
	return mergeTwoLists(left, right)
}

// 23 O(Nlog(k)), N is the total number of nodes, k is the deep of recursion
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// MergeKLists uses min heap O(Nlog(k))
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := &minHeap{}
	heap.Init(h)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	last := dummy
	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		next := minNode.Next
		last.Next = minNode
		last = minNode
		if next != nil {
			heap.Push(h, next)
		}
	}
	last.Next = nil

	return dummy.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
