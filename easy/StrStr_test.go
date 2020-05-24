package easy

import "testing"

func TestStrstr(t *testing.T) {
	seeds := []struct {
		input  string
		search string
		expect int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"a", "ab", -1},
		{"hello", "", 0},
		{"hello", "world", -1},
		{"hello", "ll", 2},
		{"hello llow", "ll", 2},
		{"mississippi", "issipi", -1},
	}

	for _, v := range seeds {
		result := strStr(v.input, v.search)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("Strstr passed !")
}
