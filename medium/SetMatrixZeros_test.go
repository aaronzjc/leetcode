package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestSetZeros(t *testing.T) {
	seeds := []struct {
		input  [][]int
		expect [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for _, v := range seeds {
		setZeroes(v.input)
		if !tools.IsL2IntArrayEquals(v.input, v.expect, true) {
			t.Error(v)
			t.Fatalf("failed !")
		}
	}

	t.Log("SetMatrixZeros passed !")
}
