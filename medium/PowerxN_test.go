package medium

import "testing"

func TestPowerxN(t *testing.T) {
	seeds := []struct {
		x      float64
		n      int
		expect float64
	}{
		{2.0, 10, 1024.0},
		{2.0, 0, 1.0},
		{2.0, -1, 0.5},
	}

	for _, v := range seeds {
		result := myPow(v.x, v.n)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("PowerxN passed !")
}
