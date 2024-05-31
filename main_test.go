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
