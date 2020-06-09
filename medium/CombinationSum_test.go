package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{{7}, {2, 2, 3}},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, v := range seeds {
		result := combinationSum(v.input, v.target)
		if !tools.IsL2IntArrayEquals(result, v.expect, false) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("CombinationSum passed !")
}
