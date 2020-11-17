package medium

import "strings"

// https://leetcode-cn.com/problems/simplify-path/

// Tips:
// 主要是栈的使用

func simplifyPath(path string) string {
	var stack []string
	is := strings.Split(path, "/")
	for i := 1; i < len(is); i++ {
		if is[i] == "" || is[i] == "." {
			continue
		}
		if is[i] == ".." {
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, is[i])
	}

	return "/" + strings.Join(stack, "/")
}
