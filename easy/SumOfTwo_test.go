package easy

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestSumTwo(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
		{[]int{2, 7, 11, 15}, 100, []int{}},
	}

	for _, v := range seeds {
		result := twoSum(v.input, v.target)
		if !tools.IsArrEquals(result, v.expect, false) {
			t.Error(v.input, v.target, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("SumTwo passed")
}
