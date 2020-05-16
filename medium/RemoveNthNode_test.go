package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestRemoveNthNode(t *testing.T) {
	seeds := []struct {
		input  []int
		idx    int
		expect []int
	}{
		// {[]int{1,2,3,4,5,6,7},2,[]int{1,2,3,4,5,7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{2, 3, 4, 5, 6, 7}},
	}

	for _, v := range seeds {
		l := tools.BuildLinklist(v.input)
		r := removeNthFromEnd(l, v.idx)
		result := tools.LoopLinklist(r)
		if !tools.IsIntArrEquals(result, v.expect, true) {
			t.Error(result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("RemvoeNthNode passed")
}
