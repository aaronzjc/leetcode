package tools

import (
	"testing"
)

func TestPreOrder(t *testing.T) {
	seeds := []struct {
		data       []int
		expectPre  []int
		expectIn   []int
		expectPost []int
		expectBfs  []int
		expectDfs  []int
	}{
		{
			/**
			 * 树的示例如下
			 *	     1
			 *	 (2     3)
			 *  (4 5) (# 6)
			 *	 (7 8)
			 */
			[]int{1, 2, 3, 4, 5, -1, 6, -1, -1, 7, 8, -1, -1},
			[]int{1, 2, 4, 5, 7, 8, 3, 6},
			[]int{4, 2, 7, 5, 8, 1, 3, 6},
			[]int{4, 7, 8, 5, 2, 6, 3, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{1, 2, 4, 5, 7, 8, 3, 6},
		},
		{
			[]int{},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for _, v := range seeds {
		tree := NewBinaryTree(v.data)
		pre1 := tree.PreOrderTravel()
		if !IsArrEquals(pre1, v.expectPre, true) {
			t.Error(pre1, v.expectPre)
			t.Fatal("failed !")
		}
		pre2 := tree.PreOrderTravelRecurse()
		if !IsArrEquals(pre2, v.expectPre, true) {
			t.Error(pre2, v.expectPre)
			t.Fatal("failed !")
		}
		in := tree.InOrderTravelRecurse()
		if !IsArrEquals(in, v.expectIn, true) {
			t.Error(in, v.expectIn)
			t.Fatal("failed !")
		}
		post := tree.PostOrderTravelRecurse()
		if !IsArrEquals(post, v.expectPost, true) {
			t.Error(post, v.expectPost)
			t.Fatal("failed !")
		}
		dfs := tree.Dfs()
		if !IsArrEquals(dfs, v.expectDfs, true) {
			t.Error(dfs, v.expectDfs)
			t.Fatal("failed !")
		}
		bfs := tree.Bfs()
		if !IsArrEquals(bfs, v.expectBfs, true) {
			t.Error(bfs, v.expectBfs)
			t.Fatal("failed !")
		}
	}

	t.Log("BinaryTree passed !")
}
