package medium

// https://leetcode-cn.com/problems/set-matrix-zeroes/

// Tips:
// 题目要求是用常数的空间来解决问题，没想出来。这里按照常规的方法解决。
// 另外思路就是利用格子里面的元素本身来存储状态，例如，遍历时，将这行和这列不为0的元素置为负数。但是这样不规范，因为没说格子本身不能为负数。

func setZeroes(matrix [][]int) {
	cols := make([]int, len(matrix[0]))
	rows := make([]int, len(matrix))

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if rows[i] == 1 || cols[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
