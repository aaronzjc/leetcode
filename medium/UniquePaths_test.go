package medium

import "testing"

func TestUniquePaths(t *testing.T) {
	seeds := []struct{
		m int
		n int
		expect int
	} {
		{7, 3, 28},
		{3, 2, 3},
	}

	for _, v := range seeds {
		result := uniquePaths(v.m, v.n)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("UniquePaths passed !")
}
