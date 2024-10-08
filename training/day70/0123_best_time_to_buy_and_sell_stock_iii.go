package day70

func maxProfit(prices []int) int {
    dp := make([][5]int, len(prices))
    dp[0][1] = -prices[0]
    dp[0][2] = 0
    dp[0][3] = -prices[0]
    dp[0][4] = 0

    for i := 1; i < len(prices); i++ {
        dp[i][1] = max(dp[i-1][1], -prices[i])
        dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
        dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
        dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
    }

    return dp[len(dp)-1][4]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
