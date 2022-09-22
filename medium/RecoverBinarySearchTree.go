package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode.cn/problems/recover-binary-search-tree/

func recoverTree(root *tools.BinaryTreeNode) {
	var inorder func(node *tools.BinaryTreeNode)
	eles := []*tools.BinaryTreeNode{}
	inorder = func(node *tools.BinaryTreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		eles = append(eles, node)
		inorder(node.Right)
	}
	inorder(root)
	var a, b *tools.BinaryTreeNode
	i, j := 0, len(eles)-1
	for (a == nil || b == nil) && i < j {
		if eles[i].Val > eles[i+1].Val {
			a = eles[i]
		} else {
			i++
		}
		if eles[j].Val < eles[j-1].Val {
			b = eles[j]
		} else {
			j--
		}
	}
	a.Val, b.Val = b.Val, a.Val
}
