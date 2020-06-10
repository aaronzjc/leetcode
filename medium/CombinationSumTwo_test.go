package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestCombinationSumTwo(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect [][]int
	}{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, v := range seeds {
		result := combinationSum2(v.input, v.target)
		if !tools.IsL2IntArrayEquals(result, v.expect, false) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("CombinationSumTwo passed !")
}
