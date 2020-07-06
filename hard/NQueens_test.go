package hard

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestQueens(t *testing.T) {
	seeds := []struct {
		input  int
		expect [][]string
	}{
		{
			4,
			[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	}

	for _, v := range seeds {
		result := solveNQueens(v.input)
		if !tools.IsL2StrArrayEquals(result, v.expect, true) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("NQueens passed !")
}
