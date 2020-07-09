package medium

import "testing"

func TestJumpGame(t *testing.T) {
	cases := []struct {
		input  []int
		expect bool
	}{
		{[]int{0}, true},
		{[]int{1}, true},
		{[]int{1, 2, 0, 1}, true},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{1, 1, 0, 1}, false},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, true},
	}

	for _, v := range cases {
		result := canJump(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("Failed !")
		}
	}

	t.Log("JumpGame passed !")
}
