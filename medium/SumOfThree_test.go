package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestThreeSum(t *testing.T) {
	seeds := []struct {
		input  []int
		expect [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
	}

	for _, v := range seeds {
		result := threeSum(v.input)

		if !tools.IsL2IntArrayEquals(result, v.expect, false) {
			t.Error(v.input, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("passed !")
}
