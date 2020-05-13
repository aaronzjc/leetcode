package medium

import "testing"

func TestIntegerToRoman(t *testing.T) {
	seeds := []struct {
		input  int
		expect string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{41, "XLI"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, v := range seeds {
		result := intToRoman(v.input)
		if result != v.expect {
			t.Log(v.input, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("IntegerToRoman passed !")
}
