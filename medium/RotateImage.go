package medium

// https://leetcode-cn.com/problems/rotate-image/

// TIPS:
// 先将矩阵转置，然后按照中间列，进行对称交换

func rotate(matrix [][]int) {
	lm := len(matrix)

	for i := 0; i < lm; i++ {
		for j := i; j < lm; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	for i := 0; i < lm; i++ {
		for j := 0; j < lm/2; j++ {
			matrix[i][j], matrix[i][lm-j-1] = matrix[i][lm-j-1], matrix[i][j]
		}
	}
}
