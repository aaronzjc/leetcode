package medium

import "testing"

func TestZigZagConvertion(t *testing.T) {
	seeds := []struct {
		input  string
		rows   int
		expect string
	}{
		{input: "", rows: 1, expect: ""},
		{input: "A", rows: 2, expect: "A"},
		{input: "AB", rows: 2, expect: "AB"},
		{input: "LEETCODEISHIRING", rows: 3, expect: "LCIRETOESIIGEDHN"},
	}

	for _, v := range seeds {
		result := zigZagConvertion(v.input, v.rows)
		if result != v.expect {
			t.Error(v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("ZigZagConvertion passed !")
}
