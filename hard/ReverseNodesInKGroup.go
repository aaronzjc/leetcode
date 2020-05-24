package hard

import (
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// Tips:
// 硬解。将链表根据k组，首先看是否有k个元素，如果没有，退出；如果有k个元素，则参考反转链表，将k个元素反转。然后拼接首尾。
// 主要是边界条件判断，链表的边界很烦。

func reverseKGroup(head *tools.ListNode, k int) (res *tools.ListNode) {
	check := func(node *tools.ListNode, k int) (match bool, first *tools.ListNode, last *tools.ListNode) {
		if node == nil {
			return
		}
		first = node
		last = node
		for node != nil && k > 1 {
			k--
			node = node.Next
			last = node
		}
		if k == 1 && node != nil {
			match = true
		}
		return
	}
	reverse := func(head *tools.ListNode, k int) (start *tools.ListNode, end *tools.ListNode) {
		var pre, tmp *tools.ListNode
		start = head
		end = head
		i := k
		for i > 1 {
			tmp = start.Next
			start.Next = pre
			pre = start
			start = tmp
			i--
		}
		start.Next = pre
		return
	}

	if head == nil {
		return
	}
	var start, end, ln, pre *tools.ListNode
	counter := 0
	res = head
	for {
		match, first, last := check(head, k)
		if !match {
			break
		}
		ln = last.Next
		start, end = reverse(first, k)
		if counter == 0 {
			res = start
		}
		if pre != nil {
			pre.Next = start
		}
		pre = end
		end.Next = ln
		head = end.Next
		counter++
	}
	return
}
