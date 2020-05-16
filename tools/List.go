package tools

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinklist(data []int) (res *ListNode) {
	var p *ListNode
	for k, v := range data {
		n := &ListNode{
			Val:  v,
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

func LoopLinklist(node *ListNode) (res []int) {
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return
}
