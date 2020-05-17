package medium

// https://leetcode-cn.com/problems/generate-parentheses/
// 关键词： 回溯；递归。

func genPath(left int, right int, path string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, path)
		return
	}

	if left > 0 {
		genPath(left-1, right, path+"(", res)
	}
	if left < right {
		genPath(left, right-1, path+")", res)
	}
}

func generateParenthesis(n int) (res []string) {
	genPath(n, n, "", &res)
	return
}
