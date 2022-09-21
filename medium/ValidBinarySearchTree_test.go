package medium

import (
	"testing"

	"github.com/aaronzjc/leetcode/tools"
)

func TestValidBinaryTree(t *testing.T) {
	seeds := []struct {
		Input  []int
		Expect bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{5, 4, 6, -1, -1, 3, 7}, false},
	}
	for _, v := range seeds {
		tree := tools.NewBinaryTree(v.Input)
		if isValidBST(tree.Root) != v.Expect {
			t.Fatal("ValidBinaryTree failed")
		}
	}

	t.Log("passed")
}
