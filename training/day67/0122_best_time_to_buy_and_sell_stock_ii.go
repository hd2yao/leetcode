package day67

func maxProfitII(prices []int) int {
    dp := make([][2]int, len(prices))
    dp[0][0] = -prices[0]
    dp[0][1] = 0

    for i := 1; i < len(prices); i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
    }

    return dp[len(prices)-1][1]
}

func maxProfitIIOneDimension(prices []int) int {
    dp := [2]int{}
    dp[0] = -prices[0]
    dp[1] = 0

    for i := 1; i < len(prices); i++ {
        dp[0] = max(dp[0], dp[1]-prices[i])
        dp[1] = max(dp[1], dp[0]+prices[i])
    }

    return dp[1]
}
