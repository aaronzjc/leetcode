package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestSpiralMatrix(t *testing.T) {
	seeds := []struct {
		input  [][]int
		expect []int
	}{
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{
				{1},
				{2},
				{3},
			},
			[]int{1, 2, 3},
		},
		{
			[][]int{
				{1, 2, 3, 4},
			},
			[]int{1, 2, 3, 4},
		},
	}

	for _, v := range seeds {
		result := spiralOrder(v.input)
		if !tools.IsArrEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("SpiralMatrix passed !")
}
