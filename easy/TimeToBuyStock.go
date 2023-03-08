package easy

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) (res int) {
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}
		if prices[i]-buy > res {
			res = prices[i] - buy
		}
	}
	return
}
