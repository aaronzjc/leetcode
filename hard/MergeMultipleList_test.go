package hard

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	seeds := []struct {
		input  [][]int
		expect []int
	}{
		{
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}, {1, 2, 3}},
			[]int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6},
		},
		{
			[][]int{{1}},
			[]int{1},
		},
	}

	for _, v := range seeds {
		llist := []*tools.ListNode{}
		for _, vv := range v.input {
			llist = append(llist, tools.BuildLinklist(vv))
		}
		r := mergeKLists(llist)
		result := tools.LoopLinklist(r)
		if !tools.IsArrEquals(result, v.expect, true) {
			t.Error(v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MergeKList passed !")
}
