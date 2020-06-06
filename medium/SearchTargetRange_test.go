package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestSearchRange(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 7, 8, 10}, 7, []int{1, 3}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1, 2, 3}, 1, []int{0, 0}},
		{[]int{}, 1, []int{-1, -1}},
	}

	for _, v := range seeds {
		result := searchRange(v.input, v.target)
		if !tools.IsIntArrEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("SearchRange passed !")
}
