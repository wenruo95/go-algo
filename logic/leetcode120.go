package logic

// leetcode 121: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxNum, maxProfit := prices[len(prices)-1], 0
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > maxNum {
			maxNum = prices[i]
		} else if v := maxNum - prices[i]; v > maxProfit {
			maxProfit = v
		}
	}
	return maxProfit
}

func MaxProfit2(prices []int) int {
	min, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if v := prices[i] - min; v > profit {
			profit = v
		}
	}
	return profit
}

func MaxProfit3(prices []int) int {
	min, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}

// leetcode 122: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii
func StockIIMaxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	dp[0] = []int{0, -prices[0]}
	for day := 1; day < n; day++ {
		dp[day] = []int{max(dp[day-1][0], dp[day-1][1]+prices[day]), max(dp[day-1][1], dp[day-1][0]-prices[day])}
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
