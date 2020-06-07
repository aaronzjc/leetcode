package medium

// https://leetcode-cn.com/problems/valid-sudoku/

// TIPS:
// 思路很简单，但是要考虑优化内存。我这里没深入优化空间。

func isValidSudoku(board [][]byte) (res bool) {
	res = true
	row := make(map[int]map[byte]bool)
	col := make(map[int]map[byte]bool)
	box := make(map[int]map[byte]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}
			// 判断行是否存在
			if _, ok := row[i][c]; ok {
				res = false
				return
			}
			if row[i] == nil {
				row[i] = make(map[byte]bool)
			}
			row[i][c] = true
			// 判断列是否存在
			if _, ok := col[j][c]; ok {
				res = false
				return
			}
			if col[j] == nil {
				col[j] = make(map[byte]bool)
			}
			col[j][c] = true
			// 判断方格是否存在
			idx := (i/3)*10 + (j / 3)
			if _, ok := box[idx][c]; ok {
				res = false
				return
			}
			if box[idx] == nil {
				box[idx] = make(map[byte]bool)
			}
			box[idx][c] = true
		}
	}
	return
}
