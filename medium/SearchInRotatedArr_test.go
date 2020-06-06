package medium

import "testing"

func TestSearchInRotatedArr(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 6, 2},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
		{[]int{9, 10, 1, 2, 3, 4, 5, 6}, 1, 2},
		{[]int{}, 5, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 3}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 1, 1},
	}

	for _, v := range seeds {
		result := search(v.input, v.target)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("SearchInRotatedArr passed !")
}
