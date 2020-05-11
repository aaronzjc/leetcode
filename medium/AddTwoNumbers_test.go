package medium

import (
	"reflect"
	"testing"
)

func genList(s []int) *ListNode {
	var l, r *ListNode
	for _, v := range s {
		tmp := &ListNode{
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

func travel(s *ListNode) (res []int) {
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
		t.Log(result, expect)
		t.Error("failed !")
	}
	t.Log("AddTwo passed")
}
