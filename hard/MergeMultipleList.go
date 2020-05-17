package hard

import (
	"github.com/aaronzjc/leetcode/easy"
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/merge-k-sorted-lists/

func MergeKLists(lists []*tools.ListNode) *tools.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var first, next *tools.ListNode
	i := 0
	for i < len(lists)-1 {
		if i == 0 {
			first = lists[i]
		}
		next = lists[i+1]
		first = easy.MergeTwoLists(first, next)
		i++
	}

	return first
}