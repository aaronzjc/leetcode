package easy

import "testing"

func TestRomanToInteger(t *testing.T) {
	seeds := []struct {
		input  string
		expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"X", 10},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, v := range seeds {
		result := romanToInt(v.input)
		if result != v.expect {
			t.Error(v.input, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("passed !")
}
