package hard

// https://leetcode-cn.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) (res int) {
	if len(s) == 0 {
		return 0
	}

	tmp := [][]int{}
	var stack []int
	var i int
	for i < len(s) {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				i++
				continue
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(tmp) == 0 {
				tmp = append(tmp, []int{last, i})
				i++
				continue
			}
			tmp = append(tmp, []int{last, i})

		MG:
			if len(tmp) >= 2 {
				t0 := tmp[len(tmp)-2]
				t1 := tmp[len(tmp)-1]
				if t0[1]+1 == t1[0] {
					tmp[len(tmp)-2] = []int{t0[0], t1[1]}
					tmp = tmp[:len(tmp)-1]
					goto MG
				} else if t0[0] > t1[0] && t0[1] < t1[1] {
					tmp[len(tmp)-2] = []int{t1[0], t1[1]}
					tmp = tmp[:len(tmp)-1]
					goto MG
				}
			}
		}
		i++
	}
	if len(tmp) == 0 {
		return 0
	}

	for _, v := range tmp {
		l := v[1] - v[0] + 1
		if l > res {
			res = l
		}
	}

	return
}
