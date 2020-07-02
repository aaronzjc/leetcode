package hard

// https://leetcode-cn.com/problems/wildcard-matching/

// TIPS:
// 参考题解，动态规划。动态规划就是一步一步推演，根据前面计算的结果，推算下一次的结果。
// 例如，计算斐波拉契数列。后面的依赖于前一次的结果，一次一次迭代下去，就得到结果了。
// 这题，必须要跳到一个更大的高度看待问题。对于`abccb`和`a*c?b`，我们要判断字符串和模式是否匹配？
// 定义dp[i][j]表示s[:i]和p[:j]匹配。只要填充这个表，然后最后输出s[len(s)][len(p)]即可。
// 问题在于怎么迭代dp状态表。对于p[j]，如果是普通字符或者?，则dp[i][j] = dp[i-1][j-1] || (s[i] == p[j] || p[j] == '?')。
// 如果p[j]是`*`，则情况复杂一些，有如下两种
// 1. 匹配0个字符，例如`a`和`a*`，此时剔除掉*，看之前是否匹配，dp[i][j] = dp[i][j-1]。
// 2. 匹配N个字符，例如`abcdef`和`a*`，当匹配`bcdef`和`*`时，可以直接忽略当前字符，只需要看前面是否匹配即可。dp[i][j] = dp[i-1][j]。
// 动态规划，不太好想上去。但是思路非常清晰，性能也比较可观。需要多训练。

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
