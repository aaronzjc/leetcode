package easy

import "testing"

func TestLenOfLastWord(t *testing.T) {
	cases := []struct {
		input  string
		expect int
	}{
		{"hello world", 5},
		{"hello world    ", 5},
		{"    ", 0},
	}

	for _, v := range cases {
		result := lengthOfLastWord(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("LenOfLastWord passed !")
}
