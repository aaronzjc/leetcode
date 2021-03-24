package easy

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) (res int) {
	buy := -1
	for i := 0; i < len(prices); i++ {
		if prices[i] < buy || buy < 0 {
			buy = prices[i]
			continue
		}
		if prices[i] > buy {
			if prices[i]-buy > res {
				res = prices[i] - buy
			}
		}
	}
	return
}
