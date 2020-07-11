package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		input  [][]int
		expect [][]int
	}{
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 0}}, [][]int{{0, 0}, {1, 4}}},
		{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}}},
		{[][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}, [][]int{{2, 4}, {5, 5}}},
	}

	for _, v := range cases {
		result := merge(v.input)
		if !tools.IsL2IntArrayEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MergeIntervals passed !")
}
