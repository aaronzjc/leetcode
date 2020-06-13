package medium

import "testing"

func TestMultiplyStrings(t *testing.T) {
	seeds := []struct {
		num1   string
		num2   string
		expect string
	}{
		{"2", "3", "6"},
		{"99", "11", "1089"},
		{"123", "456", "56088"},
		{"9133", "0", "0"},
	}

	for _, v := range seeds {
		result := multiply(v.num1, v.num2)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MultiplyStrings passed !")
}
