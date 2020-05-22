package medium

import (
	"math"
	"testing"
)

func TestStrToInt(t *testing.T) {
	seeds := []struct {
		input  string
		expect int
	}{
		{"67823", 67823},
		{"67+823", 67},
		{"67-823", 67},
		{"-67", -67},
		{"adsf", 0},
		{"67676767676767", math.MaxInt32},
		{"-6767676767", ^math.MaxInt32},
	}

	for _, v := range seeds {
		result := myAtoi(v.input)
		if result != v.expect {
			t.Error(v.input, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("StrToInt passed !")
}
