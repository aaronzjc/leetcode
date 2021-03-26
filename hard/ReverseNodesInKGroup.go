package hard

import (
	"fmt"
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// Tips:
// 暴力解法，写一个子函数，反转K个链表。然后，遍历每K个一组反转。

func reverseKGroup(head *tools.ListNode, k int) (res *tools.ListNode) {
	if k == 1 {
		res = head
		return
	}
	var reverseList func(start *tools.ListNode, end *tools.ListNode) (newStart *tools.ListNode, newEnd *tools.ListNode)
	reverseList = func(start *tools.ListNode, end *tools.ListNode) (newStart *tools.ListNode, newEnd *tools.ListNode) {
		newEnd = start
		newStart = end
		var pre, tmp *tools.ListNode
		for start != end {
			fmt.Println(start.Val)
			if pre == nil {
				pre = start
				start = start.Next
				continue
			}
			tmp = start.Next
			start.Next = pre
			pre = start
			start = tmp
		}
		start.Next = pre
		return
	}

	var s, e, ne *tools.ListNode
	i := 1
	for head != nil {
		if i%k == 1 {
			s = head
		}
		if i%k == 0 {
			e = head

			// 重置分组迭代
			i = 1
			head = head.Next

			if ne != nil {
				ne.Next = e
			}
			if res == nil {
				res = e
			}
			_, ne = reverseList(s, e)
			ne.Next = head
			continue
		}
		i++
		head = head.Next
	}
	return
}
