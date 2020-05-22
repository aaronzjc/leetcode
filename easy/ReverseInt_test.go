package easy

import "testing"

func TestReverseInt(t *testing.T) {
	seeds := []struct {
		input  int
		expect int
	}{
		{12345, 54321},
		{1<<31 - 1, 0},
		{^(1<<31 - 1), 0},
	}

	for _, v := range seeds {
		result := reverse(v.input)
		if result != v.expect {
			t.Error(result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("ReverseInt passed !")
}
