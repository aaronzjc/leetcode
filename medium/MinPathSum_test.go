package medium

import "testing"

func TestMinPathSum(t *testing.T) {
	seeds := []struct {
		grid   [][]int
		expect int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}

	for _, v := range seeds {
		result := minPathSum(v.grid)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MinPathSum passed !")
}
