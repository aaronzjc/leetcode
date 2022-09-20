package medium

import (
	"github.com/aaronzjc/leetcode/tools"
)

// https://leetcode.cn/problems/unique-binary-search-trees-ii/

func generateTrees(n int) []*tools.BinaryTree {
	var buildTree func(start, end int) []*tools.BinaryTreeNode
	buildTree = func(start, end int) []*tools.BinaryTreeNode {
		if start > end {
			return []*tools.BinaryTreeNode{nil}
		}
		trees := []*tools.BinaryTreeNode{}
		for i := start; i <= end; i++ {
			// 然后构造左子树
			leftTrees := buildTree(start, i-1)
			// 然后构造右子树
			rightTrees := buildTree(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					node := &tools.BinaryTreeNode{
						Val: i,
					}
					node.Left = left
					node.Right = right
					trees = append(trees, node)
				}
			}
		}
		return trees
	}
	var res []*tools.BinaryTree
	treeRoots := buildTree(1, n)
	for _, v := range treeRoots {
		res = append(res, &tools.BinaryTree{Root: v})
	}
	return res
}
