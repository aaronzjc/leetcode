package medium

import (
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// Tips: 思路是不管细节，先写出核心的交换逻辑，也就是29-31行。然后，再处理边界问题，再处理链接问题。

func SwapPairs(head *tools.ListNode) (rs *tools.ListNode) {
	if head == nil {
		return
	}
	rs = head
	if head.Next == nil {
		return
	}
	rs = head.Next
	var pre *tools.ListNode
	for {
		if head == nil {
			break
		}
		if head.Next == nil {
			break
		}

		i, j := head, head.Next
		i.Next = j.Next
		j.Next = i
		if pre != nil {
			pre.Next = j
		}

		pre = i
		head = i.Next
	}

	return
}
