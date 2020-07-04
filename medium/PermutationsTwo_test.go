package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestPermuteTwo(t *testing.T) {
	seeds := []struct {
		input  []int
		expect [][]int
	}{
		{
			[]int{1, 1, 2},
			[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}

	for _, v := range seeds {
		result := permuteUnique(v.input)
		if !tools.IsL2IntArrayEquals(v.expect, result, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("PermuteTwo passed !")
}
