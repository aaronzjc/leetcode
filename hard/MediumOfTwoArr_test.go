package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	expect := 2.5

	result := findMedianSortedArrays(input1,input2)

	if result != expect {
		t.Fatalf("failed !")
	}

	t.Log("FindMedianSortedArrays passed !")
}