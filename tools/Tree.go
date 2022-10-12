package tools

// Tree 树
type Tree struct {
	Length int
	Root   *TreeNode
}

// TreeNode 树节点
type TreeNode struct {
	Val      int
	Children []*TreeNode
}

// BuildTreeFromArr 根据数组生成树
func BuildTreeFromArr(data map[int][]int, keys []int) (tr *Tree) {
	var ok bool
	var node *TreeNode
	nodeMap := make(map[int]*TreeNode)
	for _, k := range keys {
		v := data[k]
		node, ok = nodeMap[k]
		if !ok && tr == nil {
			tr = &Tree{
				Length: 0,
				Root: &TreeNode{
					Val:      k,
					Children: nil,
				},
			}
			tr.Length++
			nodeMap[k] = tr.Root
			node = tr.Root
		}
		for _, vv := range v {
			tmp := &TreeNode{
				Val:      vv,
				Children: nil,
			}
			tr.Length++
			node.Children = append(node.Children, tmp)
			nodeMap[vv] = tmp
		}
	}
	return
}

// DFSTravel 深度优先遍历-非递归
func DFSTravel(t *Tree) (res []int) {
	var stack []*TreeNode
	stack = append(stack, t.Root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return
}

// DFSRecurseTravel 深度优先遍历-递归
func DFSRecurseTravel(t *Tree) (res []int) {
	var travel func(node *TreeNode, collect *[]int)
	travel = func(node *TreeNode, collec *[]int) {
		*collec = append(*collec, node.Val)
		if len(node.Children) == 0 {
			return
		}
		for _, child := range node.Children {
			travel(child, collec)
		}
	}

	travel(t.Root, &res)
	return
}

// BFSTravel 广度优先遍历-非递归
func BFSTravel(t *Tree) (res []int) {
	var queue []*TreeNode
	queue = append(queue, t.Root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		queue = append(queue, node.Children...)
		res = append(res, node.Val)
	}
	return
}
