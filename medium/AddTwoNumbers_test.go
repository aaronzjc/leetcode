package medium

import (
	"github.com/aaronzjc/leetcode/tools"
	"reflect"
	"testing"
)

func genList(s []int) *tools.ListNode {
	var l, r *tools.ListNode
	for _, v := range s {
		tmp := &tools.ListNode{
			Val:  v,
			Next: nil,
		}
		if l == nil {
			l = tmp
			r = tmp
			continue
		}
		l.Next = tmp
		l = l.Next
	}
	return r
}

func travel(s *tools.ListNode) (res []int) {
	for s != nil {
		res = append(res, s.Val)
		s = s.Next
	}
	return
}

func TestAddTwoNumbers(t *testing.T) {
	g1 := genList([]int{2, 4, 3})
	g2 := genList([]int{5, 6, 4})
	expect := []int{7, 0, 8}

	// exec
	result := travel(addTwoNumbers(g1, g2))

	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("failed !")
	}
	t.Log("AddTwo passed")
}
