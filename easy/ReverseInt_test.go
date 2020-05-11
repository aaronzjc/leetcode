package easy

import "testing"

func TestReverseInt(t *testing.T) {
	input := 12345
	expect := 54321

	result := reverse(input)

	if result != expect {
		t.Error("failed !")
	}

	t.Log("ReverseInt passed !")
}