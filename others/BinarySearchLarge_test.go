package others

import "testing"

func TestBinarySearchLarge(t *testing.T) {
	arr := []int{1, 4, 7, 10, 13, 17, 20}
	seeds := []struct {
		search int
		expect int
	}{
		{8, 3},
		{21, -1},
		{0, 0},
		{1, 1},
	}

	for _, v := range seeds {
		result := BinarySearchLarge(arr, v.search)
		if result != v.expect {
			t.Error(result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("BinarySearchLarge passed !")
}
