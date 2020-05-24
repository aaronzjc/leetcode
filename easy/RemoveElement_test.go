package easy

import "testing"

func TestRemoveElement(t *testing.T) {
	seeds := []struct {
		input  []int
		t      int
		expect int
	}{
		{[]int{}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 0, 1},
		{[]int{1, 1}, 3, 2},
		{[]int{1, 1, 2}, 1, 1},
		{[]int{1, 1, 2}, 2, 2},
		{[]int{1, 2, 2, 3, 3, 4, 5}, 3, 5},
	}

	for _, v := range seeds {
		result := removeElement(v.input, v.t)
		if result != v.expect {
			t.Error(v.input, result, v.expect, v.t)
			t.Fatalf("failed !")
		}
	}

	t.Log("RemoveElements passed !")
}
