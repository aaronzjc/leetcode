package easy

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestTopk(t *testing.T) {
	seeds := []struct {
		input  []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}

	for _, v := range seeds {
		result := Topk(v.input, v.k)
		if !tools.IsIntArrEquals(result, v.expect, false) {
			t.Log(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("Topk passed !")
}
