package solution

func maxProfitGreedy(prices []int) int {
	total := 0

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			total += prices[i] - prices[i-1]
		}
	}

	return total
}

func maxProfitDP(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	dp[0][0], dp[0][1] = 0, - prices[0]

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}
