package day53

func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2
    for i := 3; i < len(dp); i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[len(dp)-1]
}
