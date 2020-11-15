package medium

// https://leetcode-cn.com/problems/minimum-path-sum/

// Tips:
// 连续3个DP题目了。最短路径，唯一路径总数，都是分解为小问题来求解。要知道最后一个格子的值，只需要知道它相邻的两个值。
// 本题的优化点是可以利用原始数组来存储状态值，因为里面的值不会再被访问了。

func minPathSum(grid [][]int) (res int) {
	m := len(grid[0])
	n := len(grid)

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if x == 0 && y == 0 {
				grid[y][x] = grid[y][x]
				continue
			}
			if y == 0 && x > 0 {
				grid[y][x] = grid[y][x-1] + grid[y][x]
				continue
			}
			if x == 0 && y > 0 {
				grid[y][x] = grid[y-1][x] + grid[y][x]
				continue
			}
			min := grid[y-1][x]
			if min > grid[y][x-1] {
				min = grid[y][x-1]
			}
			grid[y][x] = min + grid[y][x]
		}
	}

	return grid[n-1][m-1]
}
