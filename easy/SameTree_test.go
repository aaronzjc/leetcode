package easy

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestSameTree(t *testing.T) {
	seeds := []struct {
		InputA []int
		InputB []int
		Expect bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 3, 2}, false},
		{[]int{1, 2}, []int{1, -1, 2}, false},
		{[]int{1, -1, 2}, []int{1, 2}, false},
	}
	for _, v := range seeds {
		treea := tools.NewBinaryTree(v.InputA)
		treeb := tools.NewBinaryTree(v.InputB)
		if isSameTree(treea.Root, treeb.Root) != v.Expect {
			t.Fatal("failed")
		}
	}

	t.Log("SameTree passed")
}
