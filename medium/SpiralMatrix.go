package medium

// https://leetcode-cn.com/problems/spiral-matrix/

// TIPS:
// 这题，一圈一圈的遍历就好了。

func spiralOrder(matrix [][]int) (res []int) {
	if len(matrix) <= 0 {
		return
	}
	m, n := len(matrix[0]), len(matrix)

	var walk func(int, int, int, int)
	walk = func(x int, y int, m int, n int) {
		// (0,0) -> (0,m)
		res = append(res, matrix[x][y:y+m]...)

		// (0, m) -> (n, m)
		for i := 1; i < n; i++ {
			res = append(res, matrix[x+i][y+m-1])
		}
		// (n,m) -> (n,0)
		if n > 1 {
			for i := m - 2; i >= 0; i-- {
				res = append(res, matrix[x+n-1][y+i])
			}
		}
		// (n,m) -> (0,0)
		if m > 1 {
			for i := n - 2; i > 0; i-- {
				res = append(res, matrix[x+i][y])
			}
		}

		m = m - 2
		n = n - 2
		if m <= 0 || n <= 0 {
			return
		}
		walk(x+1, y+1, m, n)
	}
	walk(0, 0, m, n)
	return
}
