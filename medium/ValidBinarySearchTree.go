package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *tools.BinaryTreeNode) bool {
	var isValid func(r *tools.BinaryTreeNode, min, max int) bool
	isValid = func(r *tools.BinaryTreeNode, min, max int) bool {
		if r == nil {
			return true
		}
		if r.Val > min && r.Val < max {
			return isValid(r.Left, min, r.Val) && isValid(r.Right, r.Val, max)
		}
		return false
	}
	return isValid(root, tools.MIN_INT, tools.MAX_INT)
}
