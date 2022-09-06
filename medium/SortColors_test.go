package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestSortColors(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	}

	for _, v := range seeds {
		sortColors(v.input)
		if !tools.IsArrEquals(v.input, v.expect, true) {
			t.Error(v)
			t.Fatalf("failed !")
		}
	}

	t.Log("SortColors passed !")
}
