package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	seeds := []struct {
		input1 []int
		input2 []int
		expect float64
	}{
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2}, []int{3, 4, 5}, 3},
		{[]int{1, 2, 3}, []int{4, 5}, 3},
		{[]int{1, 2}, []int{}, 1.5},
		{[]int{}, []int{1, 2}, 1.5},
	}

	for _, v := range seeds {
		result := findMedianSortedArrays(v.input1, v.input2)
		if result != v.expect {
			t.Error(v.input1, v.input2, result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("FindMedianSortedArrays passed !")
}
