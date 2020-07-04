package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestPermute(t *testing.T) {
	seeds := []struct {
		input  []int
		expect [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, v := range seeds {
		result := permute(v.input)
		if !tools.IsL2IntArrayEquals(v.expect, result, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("Permute passed !")
}
