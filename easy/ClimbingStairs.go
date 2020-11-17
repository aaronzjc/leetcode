package easy

// https://leetcode-cn.com/problems/climbing-stairs/

// Tips:
// 斐波拉契，动态规划

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}
