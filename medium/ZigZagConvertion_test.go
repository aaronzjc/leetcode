package medium

import "testing"

func TestZigZagConvertion(t *testing.T)  {
	input := "LEETCODEISHIRING"
	rows := 3
	expect := "LCIRETOESIIGEDHN"

	result := zigZagConvertion(input, rows)

	if result != expect {
		t.Error(expect, result)
		t.Fatalf("failed !")
	}

	t.Log("ZigZagConvertion passed !")
}