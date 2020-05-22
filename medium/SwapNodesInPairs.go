package medium

import (
	"fmt"
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

func SwapPairs(head *tools.ListNode) (rs *tools.ListNode) {
	rs = head
	if head.Next == nil {
		return
	}
	j := head
	i := head.Next
	rs = i
	for i != nil {
		j.Next = i.Next
		i.Next = j
		if j.Next == nil {
			break
		}
		j = j.Next
		if j.Next == nil {
			break
		}
		i = j.Next
		fmt.Println(j.Val, i.Val)
	}
	return
}