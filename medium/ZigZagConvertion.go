package medium

import "fmt"

// https://leetcode-cn.com/problems/zigzag-conversion/

func findNextPos(s, numRows, max int) []int {
	var next1, next2 int
	var point []int
	l := s
	point = append(point, s)
	for {
		r1 := numRows - 1 - l
		r2 := l
		next1 = 2 * r1 + s
		next2 = 2 * r2 + next1
		if next1 > max {
			break
		}
		if r1 != 0 {
			point = append(point, next1)
		}
		if next2 > max {
			break
		}
		if r2 != 0 {
			point = append(point, next2)
		}

		s = next2
	}

	return point
}

func zigZagConvertion(s string, numRows int) string {
	if s == "" || numRows == 1 {
		return s
	}
	var ps []int
	i := 0
	for i < numRows {
		t := findNextPos(i, numRows, len(s)-1)
		fmt.Println(t)
		ps = append(ps, t...)
		i++
	}
	var res string
	rs := []rune(s)
	for _, v := range ps {
		res = res + string(rs[v])
	}

	return res
}
