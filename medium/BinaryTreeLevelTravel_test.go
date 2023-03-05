package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestLevelOrder(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			[]int{3, 9, 20, -1, -1, 15, 7},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			[]int{},
			[][]int{},
		},
	}

	for _, v := range cases {
		tree := tools.NewBinaryTree(v.Input)
		result := LevelOrder(tree)
		if !tools.IsL2IntArrayEquals(result, v.Expect, true) {
			t.Error(v, result)
			t.Fatal("failed")
		}
	}
	t.Log("TestLevelOrder passed")
}
