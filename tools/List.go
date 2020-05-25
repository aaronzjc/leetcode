package tools

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildLinklist 生成单链表
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

// LoopLinklist 遍历单链表
func LoopLinklist(node *ListNode) (res []int) {
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return
}

// ReverseLinklist 反转单链表
func ReverseLinklist(node *ListNode) (start *ListNode, end *ListNode) {
	var pre, tmp *ListNode
	if node == nil {
		return
	}
	start = node
	end = node
	for start != nil {
		if start.Next == nil {
			start.Next = pre
			break
		}
		tmp = start.Next
		start.Next = pre
		pre = start
		start = tmp
	}
	return
}

// CheckLinklist 输出链表长度
func CheckLinklist(node *ListNode) (num int, first *ListNode, last *ListNode) {
	first = node
	last = node
	for node != nil {
		num++
		node = node.Next
		last = node
	}
	return
}
