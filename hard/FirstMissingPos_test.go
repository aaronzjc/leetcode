package hard

import "testing"

func TestFirstMissingPos(t *testing.T) {
	seeds := []struct {
		input  []int
		expect int
	}{
		{[]int{2}, 1},
		{[]int{1}, 2},
		{[]int{0, -1, 3, 1}, 2},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{5, 4, 3, 2, 1}, 6},
		{[]int{1, 2, 7, 4, 3}, 5},
	}

	for _, v := range seeds {
		result := firstMissingPositive(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("FirstMissingPos passed !")
}
