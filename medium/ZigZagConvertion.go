package medium

// https://leetcode-cn.com/problems/zigzag-conversion/

func zigZagConvertion(s string, numRows int) string {
	if s == "" || numRows == 1 || len(s) <= numRows {
		return s
	}
	var next1, next2, sp int
	var res []rune
	sr := []rune(s)
	slen := len(s) - 1
	i := 0
	for i < numRows && i < slen {
		sp = i
		res = append(res, sr[sp])
		for {
			r1 := numRows - 1 - i
			r2 := i
			next1 = 2*r1 + sp
			next2 = 2*r2 + next1
			if next1 > slen {
				break
			}
			if r1 != 0 {
				res = append(res, sr[next1])
			}
			if next2 > slen {
				break
			}
			if r2 != 0 {
				res = append(res, sr[next2])
			}

			sp = next2
		}
		i++
	}

	return string(res)
}
