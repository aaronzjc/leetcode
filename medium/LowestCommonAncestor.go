package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

func findPath(root, p *tools.BinaryTreeNode) (res []*tools.BinaryTreeNode) {
	node := root
	for node.Val != p.Val {
		res = append(res, node)
		// 如果节点的值小于查找的值，说明往右子树查找
		if node.Val < p.Val {
			node = node.Right
			continue
		}
		if node.Val > p.Val {
			node = node.Left
			continue
		}
	}
	res = append(res, node)
	return
}

func lowestCommonAncestor(root, p, q *tools.BinaryTreeNode) *tools.BinaryTreeNode {
	pathP := findPath(root, p)
	pathQ := findPath(root, q)

	var ancestor *tools.BinaryTreeNode
	for i := 0; i < len(pathQ) && i < len(pathP); i++ {
		if pathP[i].Val == pathQ[i].Val {
			ancestor = pathP[i]
		} else {
			break
		}
	}
	return ancestor
}
