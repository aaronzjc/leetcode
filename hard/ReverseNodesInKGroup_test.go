package hard

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestReverseNodesInKGroup(t *testing.T)  {
	seeds := []struct{
		input []int
		k int
		expect []int
	} {
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 1, []int{1, 2}},
		{[]int{1, 2}, 2, []int{2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{3, 2, 1, 6, 5, 4, 9, 8, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{3, 2, 1, 6, 5, 4, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{3, 2, 1, 6, 5, 4, 7}},
	}

	for _, v := range seeds {
		l1 := tools.BuildLinklist(v.input)
		r1 := reverseKGroup(l1, v.k)
		result := tools.LoopLinklist(r1)
		if !tools.IsIntArrEquals(result, v.expect, true) {
			t.Error(v.input, v.k, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("ReverseNodesInKGroup passed !")
}
