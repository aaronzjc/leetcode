package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	seeds := []struct{
		input []int
		expect []int
	} {
		{[]int{1,2,3,4,5,6}, []int{2,1,4,3,6,5}},
		//{[]int{}, []int{}},
		//{[]int{1}, []int{1}},
	}

	for _, v := range seeds {
		l1 := tools.BuildLinklist(v.input)
		r1 := tools.LoopLinklist(SwapPairs(l1))
		if !tools.IsIntArrEquals(r1, v.expect, true) {
			t.Error(r1, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("SwapNodesInPairs passed !")
}
