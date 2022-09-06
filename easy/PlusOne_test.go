package easy

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestPlusOne(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}

	for _, v := range seeds {
		result := plusOne(v.input)
		if !tools.IsArrEquals(result, v.expect, true) {
			t.Error(v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("PlusOne passed !")
}
