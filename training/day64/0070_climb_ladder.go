package day64

func climbLadder(n int, m int) int {
    dp := make([]int, n+1)
    dp[0] = 1

    for i := 1; i <= n; i++ {
        for j := 1; j <= m && i >= j; j++ {
            dp[i] += dp[i-j]
        }
    }
    return dp[n]
}
