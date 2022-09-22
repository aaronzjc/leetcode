package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestRecoverBinarySearchTree(t *testing.T) {
	seeds := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{5, 6, 3, 1, 4}, []int{5, 3, 6, 1, 4}},
	}

	for _, v := range seeds {
		tree := tools.NewBinaryTree(v.Input)
		recoverTree(tree.Root)
		res := tree.Bfs()
		if !tools.IsArrEquals(res, v.Expect, true) {
			t.Fatal("RecoverBinarySearchTree failed")
		}
	}
	t.Log("passed")
}
