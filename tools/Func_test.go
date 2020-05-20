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

func TestIsStringArrayEqual(t *testing.T) {
	seeds := []struct {
		a      []string
		b      []string
		expect bool
	}{
		{[]string{"1", "2", "3"}, []string{"1", "3", "2"}, true},
		{[]string{"1", "2", "3"}, []string{"1", "3"}, false},
		{[]string{}, []string{}, true},
	}

	for _, v := range seeds {
		result := IsStringArrEquals(v.a, v.b, false)
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
		{[][]int{}, [][]int{}, false, true},
		{[][]int{{4}}, [][]int{{1, 2, 3}, {4}}, false, false},
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

func TestBinarySearchWithRange(t *testing.T) {
	seeds := []struct {
		input  []int
		start  int
		end    int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 5, 3},
		{[]int{1, 2, 3, 4, 5}, 4, 5, 4},
	}

	// 查找大于3的最小索引
	search := 3
	for _, v := range seeds {
		result := BinarySearchWithRange(v.start, v.end, func(i int) bool {
			return v.input[i] > search
		})
		if result != v.expect {
			t.Error(v.input, v.start, v.end, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("BinarySearchWithRange passed !")
}
