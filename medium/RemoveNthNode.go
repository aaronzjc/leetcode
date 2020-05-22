package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *tools.ListNode, n int) *tools.ListNode {
	var p, pre *tools.ListNode
	if n == 0 {
		return head
	}
	p = head
	var i int
	for {
		if i == n {
			pre = head
		}
		if p.Next == nil {
			break
		}
		p = p.Next
		if pre != nil {
			pre = pre.Next
		}
		i++
	}
	if i == n-1 {
		head = head.Next
	} else {
		pre.Next = pre.Next.Next
	}
	return head
}
