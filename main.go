package main

import "math"

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

// Enumerate the middle element of a palindrome
func longestPalindrome0(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		// odd length
		p, q := i-1, i+1
		for ; p >= 0 && q < len(s) && s[p] == s[q]; p,
			q = p-1, q+1 {
		}
		if p+1 <= q-1 && q-p-1 > len(ret) {
			ret = s[p+1 : q]
		}

		// even length
		p, q = i, i+1
		for ; p >= 0 && q < len(s) && s[p] == s[q]; p, q = p-1, q+1 {
		}
		if p+1 <= q-1 && q-p-1 > len(ret) {
			ret = s[p+1 : q]
		}
	}
	return ret
}

// DP solution for the problem
func longestPalindrome1(s string) string {
	n := len(s)
	start, maxLen := 0, 0

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			if length == 1 {
				dp[l][r] = true
			} else if length == 2 {
				dp[l][r] = s[l] == s[r]
			} else {
				dp[l][r] = dp[l+1][r-1] && (s[l] == s[r])
			}

			if dp[l][r] && r-l+1 > maxLen {
				start = l
				maxLen = r - l + 1
			}
		}
	}

	return s[start : start+maxLen]
}

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}

	ret := 0
	for ; x > 0; x /= 10 {
		ret = ret*10 + x%10
		if (!negative && ret > math.MaxInt32) || (negative && -ret < math.MinInt32) {
			return 0
		}
	}
	if negative {
		return -ret
	} else {
		return ret
	}
}
