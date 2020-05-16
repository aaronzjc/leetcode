package medium

import "testing"

func TestThreeSumClosest(t *testing.T) {
	seeds := []struct {
		input  []int
		target int
		expect int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{1, 1, 1, 1}, 0, 3},
		{[]int{1, 1, -1, -1, 3}, 1, 1},
		{[]int{1, 1, -1, -1, 3}, 3, 3},
	}

	for _, v := range seeds {
		result := threeSumClosest(v.input, v.target)
		if result != v.expect {
			t.Error(v.input, v.target, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("ThreeSumClosest passed !")
}
