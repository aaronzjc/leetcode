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
