package medium

// https://leetcode-cn.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid[0])
	n := len(obstacleGrid)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
				continue
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]
}
