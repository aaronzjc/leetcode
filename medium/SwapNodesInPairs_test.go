package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestSwapPairs(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
	}

	for _, v := range seeds {
		l1 := tools.BuildLinklist(v.input)
		r1 := tools.LoopLinklist(swapPairs(l1))
		if !tools.IsIntArrEquals(r1, v.expect, true) {
			t.Error(r1, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("SwapNodesInPairs passed !")
}
