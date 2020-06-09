package easy

import "testing"

func TestCountAndSay(t *testing.T) {
	seeds := []struct {
		input  int
		expect string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
	}

	for _, v := range seeds {
		result := countAndSay(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("CountAndSay passed !")
}
