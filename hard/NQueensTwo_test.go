package hard

import (
	"testing"
)

func TestTotalQueens(t *testing.T) {
	seeds := []struct {
		input  int
		expect int
	}{
		{
			4,
			2,
		},
	}

	for _, v := range seeds {
		result := totalNQueens(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("NQueensTotal passed !")
}
