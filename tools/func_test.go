package tools

import "testing"

func TestIsIntArrayEqual(t *testing.T) {
	seeds := []struct {
		a      []int
		b      []int
		expect bool
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}, true},
		{[]int{1, 2, 3}, []int{1, 3}, false},
		{[]int{}, []int{}, true},
	}

	for _, v := range seeds {
		result := IsIntArrEquals(v.a, v.b, false)
		if result != v.expect {
			t.Error(v.a, v.b, v.expect, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("passed !")
}

func TestIsL2ArrayEquals(t *testing.T) {
	seeds := []struct {
		a      [][]int
		b      [][]int
		order  bool
		expect bool
	}{
		{[][]int{{1, 2, 3}, {4}}, [][]int{{1, 3, 2}, {4}}, true, false},
		{[][]int{{1, 2, 3}, {4}}, [][]int{{1, 3, 2}, {4}}, false, true},
		{[][]int{{1, 3}, {4}}, [][]int{{1, 2, 3}, {4}}, false, false},
	}

	for _, v := range seeds {
		result := IsL2IntArrayEquals(v.a, v.b, v.order)
		if result != v.expect {
			t.Error(v.a, v.b, v.expect, result)
			t.Fatalf("failed !")
		}
	}
	t.Log("passed !")
}
