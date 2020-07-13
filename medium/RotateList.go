package medium

import (
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/rotate-list/

func rotateRight(head *tools.ListNode, k int) (res *tools.ListNode) {
	if head == nil {
		return head
	}
	var s, e, l *tools.ListNode
	var length int
	s, l = head, head
	for head.Next != nil {
		length++
		head = head.Next
	}
	if head != nil {
		length++
	}

	e = head
	k = k % length

	if k == 0 {
		return s
	}

	for i := 1; i < length - k; i++ {
		l = l.Next		
	}
	res = l.Next
	l.Next = nil
	e.Next = s
	return 	
}