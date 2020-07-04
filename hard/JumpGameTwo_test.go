package hard

import "testing"

func TestJumpGameTwo(t *testing.T) {
	seeds := []struct {
		input  []int
		expect int
	}{
		{[]int{0}, 0},
		{[]int{1, 2, 1, 1, 1}, 3},
		{[]int{2, 3, 1, 4}, 2},
		{[]int{2, 3, 1, 2, 4, 2, 3}, 3},
		{[]int{1, 1, 1}, 2},
		{[]int{4, 1, 1, 3, 1, 1, 1}, 2},
		{[]int{10, 1, 1}, 1},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}, 2},
	}

	for _, v := range seeds {
		result := jump(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("JumpTwo passed !")
}
