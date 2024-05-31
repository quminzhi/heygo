package main

func twoSum(nums []int, target int) []int {
	// map from val -> idx
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok := seen[target-nums[i]]; ok {
			return []int{idx, i}
		} else {
			seen[nums[i]] = i
		}
	}

	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	c := 0

	head := &ListNode{Val: -1, Next: nil}
	tail := head

	for p != nil || q != nil || c > 0 {
		if p != nil {
			c += p.Val
			p = p.Next
		}
		if q != nil {
			c += q.Val
			q = q.Next
		}

		node := &ListNode{Val: c % 10, Next: nil}
		c /= 10

		tail.Next = node
		tail = node
	}

	return head.Next
}

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	res := 0
	for i, j := 0, 0; j < len(s); j++ {
		// s[j] is going to be added to the seen
		for i < j {
			if _, ok := seen[s[j]]; ok {
				delete(seen, s[i])
				i++
			} else {
				break
			}
		}
		seen[s[j]] = 1
		res = max(res, j-i+1)
	}

	return res
}
