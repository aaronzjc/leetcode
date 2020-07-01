package hard

// https://leetcode-cn.com/problems/wildcard-matching/

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
			} else {
				if s[i] == p[j] || p[j] == '?' {
					dp[i+1][j+1] = dp[i][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
