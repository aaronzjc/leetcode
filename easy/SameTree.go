package easy

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode.cn/problems/same-tree/

func isSameTree(p *tools.BinaryTreeNode, q *tools.BinaryTreeNode) bool {
	var compareTree func(p, q *tools.BinaryTreeNode) bool
	compareTree = func(p, q *tools.BinaryTreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil && q != nil {
			return false
		}
		if p != nil && q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return compareTree(p.Left, q.Left) && compareTree(p.Right, q.Right)
	}
	return compareTree(p, q)
}
