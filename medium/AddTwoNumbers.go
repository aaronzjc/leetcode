package medium

// https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, tmp *ListNode
	var val, pass int
	for l1 != nil || l2 != nil {
		val = 0
		node := ListNode{
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
		tmp.Next = &ListNode{
			Val:  pass,
			Next: nil,
		}
	}
	return res
}
