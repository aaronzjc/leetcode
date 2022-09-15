package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestLowestCommonAncestor(t *testing.T) {
	treeData := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	tree := tools.NewBinaryTree(treeData)

	seeds := []struct {
		input  [2]int
		expect int
	}{
		{[2]int{2, 8}, 6},
		{[2]int{2, 4}, 2},
	}
	for _, v := range seeds {
		ancestor := lowestCommonAncestor(tree.Root, &tools.BinaryTreeNode{Val: v.input[0]}, &tools.BinaryTreeNode{Val: v.input[1]})
		if ancestor.Val != v.expect {
			t.Fatalf("failed !")
		}
	}

	t.Log("LowestCommonAncestor passed !")
}
