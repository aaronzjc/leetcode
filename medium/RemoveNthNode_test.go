package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func buildRemoveNthNodeLinkList(data []int) (res *RemoveNthListNode) {
	var p *RemoveNthListNode
	for k, v := range data {
		n := &RemoveNthListNode{
			Val: v,
			Next: nil,
		}
		if k == 0 {
			p = n
			res = p
		} else {
			p.Next = n
			p = p.Next
		}
	}
	return
}

func loopRemoveNthLinkList(node *RemoveNthListNode) (res []int) {
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return
}

func TestRemoveNthNode(t *testing.T) {
	seeds := []struct{
		input []int
		idx int
		expect []int
	} {
		// {[]int{1,2,3,4,5,6,7},2,[]int{1,2,3,4,5,7}},
		{[]int{1,2,3,4,5,6,7},7,[]int{2,3,4,5,6,7}},
	}

	for _, v := range seeds {
		l := buildRemoveNthNodeLinkList(v.input)
		r := removeNthFromEnd(l, v.idx)
		result := loopRemoveNthLinkList(r)
		if !tools.IsIntArrEquals(result, v.expect, true) {
			t.Error(result, v.expect)
			t.Fatalf("failed !")
		}
	}

	t.Log("RemvoeNthNode passed")
}