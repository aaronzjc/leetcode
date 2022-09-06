package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, v := range seeds {
		nextPermutation(v.input)
		if !tools.IsArrEquals(v.input, v.expect, true) {
			t.Error(v)
			t.Fatalf("failed !")
		}
	}

	t.Log("NextPermutation passed !")
}
