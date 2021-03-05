package others

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4, 5}, []int{1, 2, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, v := range seeds {
		BubbleSort(v.input)
		if !tools.IsIntArrEquals(v.input, v.expect, true) {
			t.Log(v)
			t.Fatal("failed !")
		}
	}

	t.Log("BubbleSort passed !")
}

func TestInsertSort(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4, 5}, []int{1, 2, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, v := range seeds {
		InsertSort(v.input)
		if !tools.IsIntArrEquals(v.input, v.expect, true) {
			t.Log(v)
			t.Fatal("failed !")
		}
	}

	t.Log("InsertSort passed !")
}

func TestQuickSort(t *testing.T) {
	seeds := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4, 5}, []int{1, 2, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, v := range seeds {
		QuickSort(v.input)
		if !tools.IsIntArrEquals(v.input, v.expect, true) {
			t.Log(v)
			t.Fatal("failed !")
		}
	}

	t.Log("QuickSort passed !")
}
