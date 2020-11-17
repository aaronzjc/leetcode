package easy

import "testing"

func TestClimbStairs(t *testing.T) {
	seeds := []struct {
		input  int
		expect int
	}{
		{2, 2},
		{3, 3},
		{44, 1134903170},
	}

	for _, v := range seeds {
		result := climbStairs(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("ClimbingStairs passed !")
}
