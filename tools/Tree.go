package tools

type Tree struct {
	Length int
	Root   *TreeNode
}

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

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

// DFS
func DFSTravel(t *Tree) (res []int) {
	var l []*TreeNode
	l = append(l, t.Root)
	for len(l) > 0 {
		node := l[0]
		res = append(res, node.Val)
		l = l[1:]
		i := len(node.Children) - 1
		for i >= 0 {
			child := node.Children[i]
			l = append([]*TreeNode{child}, l...)
			i--
		}
	}
	return
}

// DFS递归版本
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

// BFS
func BFSTravel(t *Tree) (res []int) {
	var l []*TreeNode
	l = append(l, t.Root)
	for len(l) > 0 {
		for _, node := range l {
			res = append(res, node.Val)
			l = l[1:]
			for _, child := range node.Children {
				l = append(l, child)
			}
		}
	}
	return
}
