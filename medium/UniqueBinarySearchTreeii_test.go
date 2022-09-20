package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestGenerateBinaryTreeII(t *testing.T) {
	seeds := []struct {
		Input  int
		Expect [][]int
	}{
		{
			Input:  3,
			Expect: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, s := range seeds {
		trees := generateTrees(s.Input)
		var res [][]int
		for _, v := range trees {
			res = append(res, v.InOrderTravelRecurse())
		}
		if !tools.IsL2IntArrayEquals(res, s.Expect, false) {
			t.Fatal("UniqueBinarySearchTreeII failed")
		}
	}
	t.Log("passed")
}
