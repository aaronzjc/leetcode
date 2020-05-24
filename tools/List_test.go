package tools

import (
	"fmt"
	"testing"
)

func TestLoopLinklist(t *testing.T) {
	l := &ListNode{Val: 1, Next: nil}
	head := l
	l.Next = &ListNode{Val: 3, Next: nil}
	l = l.Next
	l.Next = &ListNode{Val: 2, Next: nil}
	l = l.Next
	l.Next = &ListNode{Val: 7, Next: nil}
	l = l.Next
	l.Next = &ListNode{Val: 4, Next: nil}
	l = l.Next
	l.Next = &ListNode{Val: 5, Next: nil}

	fmt.Println(head.Next)

	expect := []int{1, 3, 2, 7, 4, 5}
	result := LoopLinklist(head)
	if !IsIntArrEquals(expect, result, true) {
		t.Error(expect, result)
		t.Fatalf("failed !")
	}

	t.Log("BuildLinklist passed !")
}

func TestBuildLinklist(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 3, 2, 7, 4, 5}, []int{1, 3, 2, 7, 4, 5}},
		{[]int{}, nil},
	}

	for _, v := range seeds {
		l := BuildLinklist(v.input)
		r := LoopLinklist(l)
		if !IsIntArrEquals(v.input, v.expect, true) {
			t.Error(v.input, r)
			t.Fatalf("failed !")
		}
	}

	t.Log("BuildLinklist passed !")
}

func TestReverseLinklist(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, v := range seeds {
		l := BuildLinklist(v.input)
		r, _ := ReverseLinklist(l)
		result := LoopLinklist(r)
		if !IsIntArrEquals(result, v.expect, true) {
			t.Error(v.input, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("ReverseLinklist passed !")
}

func TestCheckLinklist(t *testing.T) {
	seeds := []struct {
		input    []int
		expNum   int
		expFirst int
		expLast  int
	}{
		{[]int{}, 0, -1, -1},
		{[]int{1}, 1, 1, 1},
		{[]int{1, 2}, 2, 1, 2},
	}

	for _, v := range seeds {
		l := BuildLinklist(v.input)
		result, first, last := CheckLinklist(l)
		if result != v.expNum &&
			((first == nil && v.expFirst == -1) || (first != nil && v.expFirst == first.Val)) &&
			((last == nil && v.expLast == -1) || (last != nil && v.expLast == last.Val)) {
			t.Error(v, result, first, last)
			t.Fatalf("failed !")
		}
	}

	t.Log("CheckLinklist passed !")
}
