package easy

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

var (
	TopkSeeds = []struct {
		input  []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}
)

func TestTopk(t *testing.T) {
	for _, v := range TopkSeeds {
		result := Topk(v.input, v.k)
		if !tools.IsArrEquals(result, v.expect, false) {
			t.Log(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("Topk passed !")
}

func TestTopk2(t *testing.T) {
	for _, v := range TopkSeeds {
		result := Topk2(v.input, v.k)
		if !tools.IsArrEquals(result, v.expect, false) {
			t.Log(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("Topk passed !")
}
