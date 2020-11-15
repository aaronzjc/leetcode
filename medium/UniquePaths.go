package medium

import (
	"strconv"
)

// https://leetcode-cn.com/problems/unique-paths/

// Tips:
// 第一想法是用回溯法，把所有的路径都算出来，然后统计。但是显然如果棋盘太大，会出现性能问题。
// 然后，再仔细研究，发现这是一个典型的DP问题。要知道从起点到终点有多少个路径，只需要知道前一个格子到终点有多少路径，然后，累计即可。
// 很像斐波拉契数列的方法。

func uniquePaths(m int, n int) int {
	dp := make(map[string]int)
	var calArea func(int, int) int
	calArea = func(mm int, nn int) int {
		if (mm + nn - 2) <= 1 {
			return 1
		}
		var x, y, p1, p2 int
		var key string
		x = mm - 1; y = nn
		if x >= 1 && y >= 1 {
			key = strconv.Itoa(x) + "_" + strconv.Itoa(y)
			if v, ok := dp[key]; ok {
				p1 = v
			} else {
				p1 = calArea(x, y)
				dp[key] = p1
			}
		}
		x = mm; y = nn - 1
		if x >= 1 && y >= 1 {
			key = strconv.Itoa(x) + "_" + strconv.Itoa(y)
			if v, ok := dp[key]; ok {
				p2 = v
			} else {
				p2 = calArea(x, y)
				dp[key] = p2
			}
		}
		return p1 + p2
	}
	return calArea(m, n)
}