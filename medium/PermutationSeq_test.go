package medium

import "testing"

func TestGetPermutationK(t *testing.T)  {
	cases := []struct{
		n int
		k int
		expect string
	} {
		{3, 3, "213"},
		{8, 13122, "36247851"},
	}

	for _, v := range cases {
		result := getPermutation(v.n, v.k)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("GetKPermutation passed !")
}
