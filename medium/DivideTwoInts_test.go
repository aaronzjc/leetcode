package medium

import (
	"math"
	"testing"
)

func TestDivideZero(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Did panic !")
		}
	}()

	divide(10, 0)
}

func TestDivideTwoInts(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Log("none recover emitted !")
		}
	}()
	seeds := []struct {
		a      int
		b      int
		expect int
	}{
		{10, 3, 3},
		{10, 2, 5},
		{29, 2, 14},
		{7, -3, -2},
		{1, 1, 1},
		{1, 1, 1},
		{2, 3, 0},
		{-2147483648, -1, math.MaxInt32},
		{-2147483648, 1, ^math.MaxInt32},
		{-2147483648, 2, -1073741824},
	}

	for _, v := range seeds {
		result := divide(v.a, v.b)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("DivideTwoInts passed !")
}
