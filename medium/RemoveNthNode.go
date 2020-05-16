package medium

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

type RemoveNthListNode struct {
    Val int
    Next *RemoveNthListNode
 }

func removeNthFromEnd(head *RemoveNthListNode, n int) *RemoveNthListNode {
    var p, pre *RemoveNthListNode
    p = head
    var i int
    for {
        if i == n {
            pre = head
        }
        if p.Next == nil {
            break
        }
        p = p.Next
        if pre != nil {
            pre = pre.Next
        }
        i++
    }
    if i == n -1 {
        head = head.Next
    } else {
        pre.Next = pre.Next.Next
    }
    return head
}