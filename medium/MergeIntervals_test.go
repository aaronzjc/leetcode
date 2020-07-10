package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestMergeIntervals(t *testing.T)  {
	cases := []struct{
		input [][]int
		expect [][]int
	} {
		{[][]int{{1,3}, {2,6},{8,10},{15,18}}, [][]int{{1,6}, {8,10}, {15,18}}},
		{[][]int{{1,4}, {4,5}}, [][]int{{1,5}}},
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
