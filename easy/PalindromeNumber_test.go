package easy

import "testing"

func TestIsPalindromeNumber(t *testing.T) {
	seeds := []struct {
		input  int
		expect bool
	}{
		{-123, false},
		{1, true},
		{1222221, true},
		{12344321, true},
	}
	for _, v := range seeds {
		result := isPalindrome(v.input)
		if result != v.expect {
			t.Error(v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("IsPalindromeNumber passed !")
}
