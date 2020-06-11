package hard

import (
	"sort"

	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode-cn.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*tools.ListNode) (rs *tools.ListNode) {
	var nodes []*tools.ListNode
	var i int
	for _, l := range lists {
		i = 0
		for l != nil {
			i = sort.Search(len(nodes), func(i int) bool {
				return nodes[i].Val > l.Val
			})
			if i == len(nodes) || len(nodes) == 0 {
				nodes = append(nodes, &tools.ListNode{
					Val:  l.Val,
					Next: nil,
				})
			} else {
				nodes = append(nodes, nil)
				copy(nodes[i+1:], nodes[i:])
				nodes[i] = &tools.ListNode{
					Val:  l.Val,
					Next: nil,
				}
			}
			l = l.Next
		}
	}

	for k := 0; k < len(nodes); k++ {
		if k == 0 {
			rs = nodes[k]
		}
		if k == len(nodes)-1 {
			continue
		}
		nodes[k].Next = nodes[k+1]
	}

	return
}
