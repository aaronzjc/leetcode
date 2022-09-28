package tools

// BinaryTree 二叉树
type BinaryTree struct {
	Root *BinaryTreeNode
}

// BinaryTreeNode 二叉树节点
type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Val   int
}

// NewBinaryTree 根据层次遍历数组构造一个二叉树，对于没有值的节点，采用虚拟节点
func NewBinaryTree(data []int) (tree *BinaryTree) {
	var build func(*BinaryTreeNode, int)
	build = func(node *BinaryTreeNode, i int) {
		node.Val = data[i]
		leftIdx, rightIdx := i*2+1, i*2+2
		if leftIdx >= len(data) || data[leftIdx] == -1 {
			node.Left = nil
		} else {
			node.Left = &BinaryTreeNode{}
			build(node.Left, leftIdx)
		}
		if rightIdx >= len(data) || data[rightIdx] == -1 {
			node.Right = nil
		} else {
			node.Right = &BinaryTreeNode{}
			build(node.Right, i*2+2)
		}
	}
	tree = &BinaryTree{}
	if len(data) == 0 {
		return
	}
	tree.Root = new(BinaryTreeNode)
	build(tree.Root, 0)
	return
}

// PreOrderTravelRecurse 先序遍历-递归
func (t *BinaryTree) PreOrderTravelRecurse() (res []int) {
	var loop func(node *BinaryTreeNode)
	loop = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		loop(node.Left)
		loop(node.Right)
	}
	loop(t.Root)
	return
}

// PreOrderTravel 先序遍历-非递归
func (t *BinaryTree) PreOrderTravel() (res []int) {
	if t.Root == nil {
		return
	}
	var stack []*BinaryTreeNode
	stack = append(stack, t.Root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}

// InOrderTravelRecurse 中序遍历-递归
func (t *BinaryTree) InOrderTravelRecurse() (res []int) {
	var loop func(node *BinaryTreeNode)
	loop = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		loop(node.Left)
		res = append(res, node.Val)
		loop(node.Right)
	}
	loop(t.Root)
	return
}

// PostOrderTravelRecurse 后序遍历-递归
func (t *BinaryTree) PostOrderTravelRecurse() (res []int) {
	var loop func(node *BinaryTreeNode)
	loop = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		loop(node.Left)
		loop(node.Right)
		res = append(res, node.Val)
	}
	loop(t.Root)
	return
}

// Dfs Dfs遍历
func (t *BinaryTree) Dfs() (res []int) {
	if t.Root == nil {
		return
	}
	// Dfs是先进后出，所以用栈
	stack := []*BinaryTreeNode{t.Root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return
}

// Bfs Bfs遍历
func (t *BinaryTree) Bfs() (res []int) {
	if t.Root == nil {
		return
	}
	// Bfs遍历是先进先出，所以用队列
	pipe := []*BinaryTreeNode{t.Root}
	for len(pipe) > 0 {
		node := pipe[0]
		pipe = pipe[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			pipe = append(pipe, node.Left)
		}
		if node.Right != nil {
			pipe = append(pipe, node.Right)
		}
	}
	return
}
