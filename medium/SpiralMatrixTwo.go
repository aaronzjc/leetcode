package medium

// https://leetcode-cn.com/problems/spiral-matrix-ii/

func generateMatrix(n int) (res [][]int) {
	res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	counter := 0
	var walk func(int, int, int)
	walk = func(x int, y int, size int) {
		// (0, 0) -> (n, 0)
		for i := 0; i < size; i++ {
			counter++
			res[x][i+y] = counter
		}
		// (n, 0) -> (n, n)
		for i := 1; i < size; i++ {
			counter++
			res[x+i][y+size-1] = counter
		}
		// (n, n) -> (0, n)
		for i := size - 2; i >= 0; i-- {
			counter++
			res[x+size-1][y+i] = counter
		}
		// (0, n) -> (0, 0)
		for i := size - 2; i > 0; i-- {
			counter++
			res[x+i][y] = counter
		}
		n = n - 2
		if n <= 0 {
			return
		}
		walk(x+1, y+1, n)
	}

	walk(0, 0, n)
	return
}
