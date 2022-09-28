package medium

// https://leetcode-cn.com/problems/coin-change/solution/golangdong-tai-gui-hua-by-quantumer/

// Tips：
// dp[i] = min(dp[i], dp[i-coin])

// CoinChange 换硬币
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i < coin || dp[i-coin] == -1 {
				continue
			}
			count := 1 + dp[i-coin]
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}
