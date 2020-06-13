package hard

import "testing"

func TestTrappingRain(t *testing.T) {
	seeds := []struct {
		input  []int
		expect int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, v := range seeds {
		result := trap(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("TrappingRain passed !")
}
