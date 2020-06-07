package hard

// https://leetcode-cn.com/problems/sudoku-solver/

// TIPS:
// 回溯法回溯法回溯法！！

func solveSudoku(board [][]byte) {
	// 检查当前坐标能否填这个值
	isValid := func(i int, j int, c byte) bool {
		for ii := 0; ii < 9; ii++ {
			if board[i][ii] == c || board[ii][j] == c {
				return false
			}
		}

		iStart := (i / 3) * 3
		iEnd := iStart + 3
		jStart := (j / 3) * 3
		jEnd := jStart + 3
		for ii := iStart; ii < iEnd; ii++ {
			for jj := jStart; jj < jEnd; jj++ {
				if board[ii][jj] == c {
					return false
				}
			}
		}

		return true
	}

	var resolve func([][]byte, int, int) bool
	resolve = func(board [][]byte, i int, j int) bool {
		ii := i + j / 9
		jj := j % 9
		for ; ii < 9; ii++ {
			for ; jj < 9; jj++ {
				// 这个位置有数字，跳过
				if board[ii][jj] != '.' {
					continue
				}

				for k := 0; k < 9; k++ {
					ck := byte('1' + k)
					// 如果这个位置填入这个数字有效，则填下一个
					if !isValid(ii, jj, ck) {
						continue
					}
					board[ii][jj] = ck
					ret := resolve(board, ii, jj+1)
					if ret {
						return true
					}
					board[ii][jj] = '.'
				}
				return false
			}
			jj = 0
		}

		return true
	}

	resolve(board, 0, 0)
}
