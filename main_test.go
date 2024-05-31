package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "t1: 0 + 0", args: args{l1: &ListNode{Val: 0},
			l2: &ListNode{Val: 0}},
			want: &ListNode{Val: 0, Next: nil}},
		{name: "t1: 342 + 465 = 807", args: args{l1: &ListNode{Val: 2,
			Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4,
				Next: nil}}}}, want: &ListNode{Val: 7, Next: &ListNode{Val: 0,
			Next: &ListNode{Val: 8, Next: nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{s: "abcabcbb"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "t1", args: args{nums: []int{1, 2, 3}, target: 5}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome0(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{s: "babad"}, want: "bab"},
		{name: "t2", args: args{s: "baab"}, want: "baab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome0(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{s: "babad"}, want: "bab"},
		{name: "t2", args: args{s: "baab"}, want: "baab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{x: 5}, want: 5},
		{name: "t2", args: args{x: 1234}, want: 4321},
		{name: "t3", args: args{x: -1234}, want: -4321},
		{name: "t4", args: args{x: 120}, want: 21},
		{name: "t5", args: args{x: 1534236469}, want: 0},
		{name: "t6", args: args{x: -2147483648}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
