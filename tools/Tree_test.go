package tools

import "testing"

func TestBuildTreeFromArr(t *testing.T) {
	// 手动生成一个书
	tr := &Tree{Length: 0, Root: nil}
	tr.Root = &TreeNode{Val: 1, Children: []*TreeNode{
		{
			Val:      2,
			Children: nil,
		},
		{
			Val: 3,
			Children: []*TreeNode{
				{5, nil},
				{6, nil},
			},
		},
		{
			Val: 4,
			Children: []*TreeNode{
				{7, nil},
			},
		},
	}}

	// 自动生成一个树
	input := map[int][]int{
		1: {2, 3, 4},
		2: nil,
		3: {5, 6},
		4: {7},
	}
	tt := BuildTreeFromArr(input, []int{1, 2, 3, 4})

	res1 := BFSTravel(tr)
	res2 := BFSTravel(tt)

	if !IsArrEquals(res1, res2, true) {
		t.Error(res1, res2)
		t.Fatalf("failed !")
	}

	t.Log("BuildTreeFromArrpassed !")
}

func TestDFSTravel(t *testing.T) {
	seeds := []struct {
		tree   map[int][]int
		keys   []int
		expect []int
	}{
		{map[int][]int{1: {2, 3, 4}, 3: {5, 6}, 4: {7}}, []int{1, 3, 4}, []int{1, 2, 3, 5, 6, 4, 7}},
		{map[int][]int{1: {2, 3}, 2: {4, 5, 6}, 3: {7}}, []int{1, 2, 3}, []int{1, 2, 4, 5, 6, 3, 7}},
	}

	for _, v := range seeds {
		tr := BuildTreeFromArr(v.tree, v.keys)
		trs := DFSTravel(tr)
		if !IsArrEquals(trs, v.expect, true) {
			t.Error(v.tree, trs, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("DFSTravel passed !")
}

func TestDFSRecurseTravel(t *testing.T) {
	seeds := []struct {
		tree   map[int][]int
		keys   []int
		expect []int
	}{
		{map[int][]int{1: {2, 3, 4}, 3: {5, 6}, 4: {7}}, []int{1, 3, 4}, []int{1, 2, 3, 5, 6, 4, 7}},
		{map[int][]int{1: {2, 3}, 2: {4, 5, 6}, 3: {7}}, []int{1, 2, 3}, []int{1, 2, 4, 5, 6, 3, 7}},
	}

	for _, v := range seeds {
		tr := BuildTreeFromArr(v.tree, v.keys)
		trs := DFSRecurseTravel(tr)
		if !IsArrEquals(trs, v.expect, true) {
			t.Error(v.tree, trs, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("DFSTravel passed !")
}

func TestBFSTravel(t *testing.T) {
	seeds := []struct {
		tree   map[int][]int
		keys   []int
		expect []int
	}{
		{map[int][]int{1: {2, 3, 4}, 3: {5, 6}, 4: {7}}, []int{1, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7}},
		{map[int][]int{1: {2, 3}, 2: {4, 5, 6}, 3: {7}}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, v := range seeds {
		tr := BuildTreeFromArr(v.tree, v.keys)
		trs := BFSTravel(tr)
		if !IsArrEquals(trs, v.expect, true) {
			t.Error(trs, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("BFSTravel passed !")
}
