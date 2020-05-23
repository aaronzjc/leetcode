package hard

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// Tips:

func reverseKGroup(head *tools.ListNode, k int) (res *tools.ListNode) {
	counter := 0
	check := func(node *tools.ListNode, k int) (match bool, first *tools.ListNode, last *tools.ListNode) {
		first = node
		for node.Next != nil || k > 1 {
			k--
			node = node.Next
		}
		if k == 1 && node.Next == nil {
			match = true
		}
		last = node
		return
	}
	for {
		match, first, last := check(head, k)
		if !match {
			break
		}
		counter++
	}
	return
}