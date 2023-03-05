package medium

import "github.com/aaronzjc/leetcode/tools"

// https://leetcode.cn/problems/binary-tree-level-order-traversal/submissions/

func LevelOrder(root *tools.BinaryTree) (res [][]int) {
	if root.Root == nil {
		return
	}
	queue := []*tools.BinaryTreeNode{}
	queue = append(queue, root.Root)
	var tmp []int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}
