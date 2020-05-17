package easy

import "github.com/aaronzjc/leetcode/tools"

func MergeTwoLists(l1 *tools.ListNode, l2 *tools.ListNode) (r *tools.ListNode) {
	var rs *tools.ListNode
	for l1 != nil || l2 != nil {
		if rs == nil {
			rs = new(tools.ListNode)
			r = rs
		} else {
			rs.Next = new(tools.ListNode)
			rs = rs.Next
		}
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				rs.Val = l1.Val
				l1 = l1.Next
			} else {
				rs.Val = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			rs.Val = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			rs.Val = l2.Val
			l2 = l2.Next
		}
	}

	return
}
