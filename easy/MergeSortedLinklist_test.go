package easy

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestMergeSortedLinklist(t *testing.T) {
	seeds := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{0}, []int{0}},
	}

	for _, v := range seeds {
		la := tools.BuildLinklist(v.a)
		lb := tools.BuildLinklist(v.b)
		r := mergeTwoLists(la, lb)
		result := tools.LoopLinklist(r)
		if !tools.IsArrEquals(result, v.expect, true) {
			t.Error(v.a, v.b, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MergeTwoList passed !")
}
