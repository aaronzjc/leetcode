package easy

import "testing"

func TestMaxSubArray(t *testing.T) {
	seeds := []struct{
		input []int
		expect int
	} {
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
	}

	for _, v := range seeds {
		result := maxSubArray(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("MaximumSubarr passed !")
}
