package medium

import "testing"

func TestStrToInt(t *testing.T) {
	input := "67823"
	expect := 67823

	result := myAtoi(input)

	if result != expect {
		t.Error("failed !")
	}

	t.Log("StrToInt passed !")
}