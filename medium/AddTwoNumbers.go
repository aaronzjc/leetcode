package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode-cn.com/problems/add-two-numbers/

func addTwoNumbers(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {
	var res, tmp *tools.ListNode
	var val, pass int
	for l1 != nil || l2 != nil {
		val = 0
		node := tools.ListNode{
			Val:  val,
			Next: nil,
		}
		if res == nil {
			tmp = &node
			res = &node
		} else {
			tmp.Next = &node
			tmp = tmp.Next
		}
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		node.Val = (val + pass) % 10
		pass = (val + pass) / 10
	}
	if pass != 0 {
		tmp.Next = &tools.ListNode{
			Val:  pass,
			Next: nil,
		}
	}
	return res
}
