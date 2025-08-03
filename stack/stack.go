package stack

import "math"

// 20
func isValid(s string) bool {
	stk := make([]rune, 0)
	for _, ch := range s {
		switch ch {
		case '(':
			stk = append(stk, ch)
		case '{':
			stk = append(stk, ch)
		case '[':
			stk = append(stk, ch)
		case ')':
			if len(stk) == 0 || stk[len(stk)-1] != '(' {
				return false
			}
			stk = stk[:len(stk)-1]
		case '}':
			if len(stk) == 0 || stk[len(stk)-1] != '{' {
				return false
			}
			stk = stk[:len(stk)-1]
		case ']':
			if len(stk) == 0 || stk[len(stk)-1] != '[' {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}

// 155
// f[i] is the minimum value of first i elements in the stack
type MinStack struct {
	stk []int
	f   []int
}

func Constructor() MinStack {
	s := MinStack{
		stk: make([]int, 0),
		f:   make([]int, 1), // 1-based
	}
	s.f[0] = math.MaxInt32
	return s
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	minInt := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	this.f = append(this.f, minInt(this.f[len(this.stk)-1], val))
}

func (this *MinStack) Pop() {
	if len(this.stk) == 0 {
		return
	}
	this.stk = this.stk[:len(this.stk)-1]
	this.f = this.f[:len(this.f)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.f[len(this.stk)]
}

// 496
// nums1 is a subset of nums2
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	candidates := make([]int, 0)
	nextGreater := make([]int, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		e := nums2[i]
		for len(candidates) != 0 && e >= candidates[len(candidates)-1] {
			candidates = candidates[:len(candidates)-1]
		}
		if len(candidates) == 0 {
			nextGreater[i] = -1
		} else {
			nextGreater[i] = candidates[len(candidates)-1]
		}
		candidates = append(candidates, e)
	}

	index := make(map[int]int)
	for i, v := range nums2 {
		index[v] = i
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = nextGreater[index[v]]
	}
	return res
}
