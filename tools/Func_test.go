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

func TestPermutation(t *testing.T) {
	seeds := []struct {
		input  []string
		expect []string
	}{
		{[]string{"hello", "world"}, []string{"helloworld", "worldhello"}},
		{[]string{"a", "b", "c"}, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{[]string{"hello"}, []string{"hello"}},
		{[]string{}, []string{}},
	}

	for _, v := range seeds {
		result := Permutation(v.input)
		if !IsStringArrEquals(v.expect, result, false) {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("Permutation passed")
}
