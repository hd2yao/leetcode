package day64

import "math"

// 二维 dp
func numSquares(n int) int {
    dp := make([][]int, 101)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, n+1)
    }

    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        dp[0][i] = math.MaxInt32
    }
    for i := 1; i <= 100; i++ {
        for j := 0; j <= n; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= i*i {
                dp[i][j] = min(dp[i][j], dp[i][j-i*i]+1)
            }
        }
    }
    return dp[100][n]
}

// 一维 dp
func numSquaresOneDimension(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = math.MaxInt32
    }
    for i := 1; i <= 100; i++ {
        for j := 0; j <= n; j++ {
            dp[j] = dp[j]
            if j >= i*i {
                dp[j] = min(dp[j], dp[j-i*i]+1)
            }
        }
    }
    return dp[n]
}
