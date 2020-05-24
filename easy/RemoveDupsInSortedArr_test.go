package easy

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	seeds := []struct {
		input  []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1, 2, 3}, 3},
		{[]int{1, 2, 2, 3, 3, 4, 5}, 5},
		{[]int{1, 2, 3, 4, 5, 5}, 5},
	}

	for _, v := range seeds {
		result := removeDuplicates(v.input)
		if result != v.expect {
			t.Error(v.input, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("RemoveDuplicates passed !")
}
