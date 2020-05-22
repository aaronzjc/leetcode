package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	seeds := []struct {
		input1 []int
		input2 []int
		expect []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{1, 2, 3}, []int{4, 5}, []int{5, 7, 3}},
		{[]int{2, 3}, []int{8, 7}, []int{0, 1, 1}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, v := range seeds {
		l1 := tools.BuildLinklist(v.input1)
		l2 := tools.BuildLinklist(v.input2)
		r1 := addTwoNumbers(l1, l2)
		result := tools.LoopLinklist(r1)
		if !tools.IsIntArrEquals(result, v.expect, true) {
			t.Error(v.input1, v.input2, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("AddTwo passed")
}
