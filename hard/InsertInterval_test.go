package hard

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestInsertInterval(t *testing.T) {
	cases := []struct {
		input  [][]int
		insert []int
		expect [][]int
	}{
		{[][]int{{1, 8}}, []int{2, 6}, [][]int{{1, 8}}},
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, v := range cases {
		result := insert(v.input, v.insert)
		if !tools.IsL2IntArrayEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("InsertInterval passed !")
}
