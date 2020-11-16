package easy

import "testing"

func TestMySqrtx(t *testing.T) {
	seeds := []struct {
		input  int
		expect int
	}{
		{4, 2},
		{8, 2},
		{1, 1},
		{2, 1},
		{4, 2},
		{5, 2},
	}

	for _, v := range seeds {
		result := mySqrt(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MySqrtx passed !")
}
