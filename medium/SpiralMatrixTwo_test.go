package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	cases := []struct {
		input  int
		expect [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, v := range cases {
		result := generateMatrix(v.input)
		if !tools.IsL2IntArrayEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("GenerateMatrix passed !")
}
