package easy

// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	var stack []byte

	for _, v := range []byte(s) {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			c := stack[len(stack)-1]
			if v == ')' && c != '(' {
				return false
			}
			if v == ']' && c != '[' {
				return false
			}
			if v == '}' && c != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
