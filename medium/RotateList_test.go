package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		input []int
		k int
		expect []int
	} {
		{[]int{1,2,3,4,5}, 2, []int{4,5,1,2,3}},
		{[]int{1,2,3,4,5}, 0, []int{1,2,3, 4, 5}},
		{[]int{0,1,2}, 4, []int{2,0,1}},
		{[]int{}, 0, []int{}},
	}

	for _, v := range cases {
		link := tools.BuildLinklist(v.input)
		result := rotateRight(link, v.k)
		reList := tools.LoopLinklist(result)
		if !tools.IsIntArrEquals(reList, v.expect, true) {
			t.Error(v, reList)
			t.Fatalf("failed !")
		}
	}

	t.Log("RotateRight passed !")
}