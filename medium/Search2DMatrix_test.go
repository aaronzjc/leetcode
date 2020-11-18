package medium

import "testing"

func TestSearchMatrix(t *testing.T) {
	seeds := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false},
		{[][]int{}, 0, false},
		{[][]int{{}}, 0, false},
		{[][]int{{1}}, 2, false},
		{[][]int{{3}}, 2, false},
	}

	for _, v := range seeds {
		result := searchMatrix(v.matrix, v.target)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("Search2DMatrix passed !")
}
