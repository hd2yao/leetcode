package day71

func maxProfit(prices []int) int {
    dp := make([][4]int, len(prices))
    dp[0][0], dp[0][1], dp[0][2], dp[0][3] = 0, -prices[0], 0, 0

    for i := 1; i < len(prices); i++ {
        dp[i][0] = dp[i-1][2]
        dp[i][1] = max(dp[i-1][1], max(dp[i-1][0]-prices[i], dp[i-1][3]-prices[i]))
        dp[i][2] = dp[i-1][1] + prices[i]
        dp[i][3] = max(dp[i-1][0], dp[i-1][3])
    }

    return max(dp[len(prices)-1][3], max(dp[len(prices)-1][0], dp[len(prices)-1][2]))
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
