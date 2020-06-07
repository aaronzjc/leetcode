package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, v := range seeds {
		result := searchInsert(v.input, v.target)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("SearchInsertPos passed !")
}
