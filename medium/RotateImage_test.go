package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestRotateImage(t *testing.T) {
	seeds := []struct{
		input [][]int
		expect [][]int
	} {
		{
			[][]int{
				{1,2,3},
				{4,5,6},
				{7,8,9},
			},
			[][]int{
				{7,4,1},
				{8,5,2},
				{9,6,3},
			},
		},
	}

	for _, v := range seeds {
		rotate(v.input)
		if !tools.IsL2IntArrayEquals(v.input, v.expect, true) {
			t.Error(v)
			t.Fatalf("failed !")
		}
	}

	t.Log("RotateImage passed !")
}