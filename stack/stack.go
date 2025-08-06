package stack

import (
	"math"
	"unicode"
)

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

// 227
// "(3+2)*2"
// Refer to inorder traversal non recursive implementation
// 栈顶操作符的优先级若大于等于当前元素（左子树），则计算
// 若小于，当前操作符属于栈顶元素的右子树，则入栈
// 若遇到右括号，栈内左括号与右括号之间的表达式优先级最高（右括号后的操作符的左子树），计算
func calculate(s string) int {
	valStack := make([]int, 0)
	opsStack := make([]rune, 0)

	precedence := make(map[rune]int)
	precedence['+'] = 0
	precedence['-'] = 0
	precedence['*'] = 1
	precedence['/'] = 1

	eval := func() {
		ops := opsStack[len(opsStack)-1]
		opsStack = opsStack[:len(opsStack)-1]
		b := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		a := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		switch ops {
		case '+':
			valStack = append(valStack, a+b)
		case '-':
			valStack = append(valStack, a-b)
		case '*':
			valStack = append(valStack, a*b)
		case '/':
			valStack = append(valStack, a/b)
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			j := i
			for j < len(s) && s[j] == ' ' {
				j++
			}
			i = j - 1
		} else if unicode.IsDigit(rune(s[i])) {
			j, num := i, 0
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				num = num*10 + int(s[j]-'0')
				j++
			}
			valStack = append(valStack, num)
			i = j - 1
		} else if s[i] == '(' {
			opsStack = append(opsStack, '(')
		} else if s[i] == ')' {
			for len(opsStack) > 0 && opsStack[len(opsStack)-1] == '(' {
				eval()
			}
			opsStack = opsStack[:len(opsStack)-1] // pop (
		} else {
			// ops
			for len(opsStack) > 0 && precedence[rune(opsStack[len(
				opsStack)-1])] >= precedence[rune(s[i])] {
				eval()
			}
			opsStack = append(opsStack, rune(s[i]))
		}
	}
	for len(opsStack) > 0 {
		eval()
	}

	return valStack[len(valStack)-1]
}

//
