package day64

import "math"

// 二维 dp
func coinChange(coins []int, amount int) int {
    dp := make([][]int, len(coins)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, amount+1)
    }
    // 初始化
    dp[0][0] = 0
    for i := 1; i <= amount; i++ {
        dp[0][i] = math.MaxInt32
    }

    for i := 1; i <= len(coins); i++ {
        for j := 0; j <= amount; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= coins[i-1] {
                dp[i][j] = min(dp[i][j], dp[i][j-coins[i-1]]+1)
            }
        }
    }
    if dp[len(coins)][amount] == math.MaxInt32 {
        return -1
    }

    return dp[len(coins)][amount]
}

// 一维 dp
func coinChangeOneDimension(coins []int, amount int) int {
    dp := make([]int, amount+1)
    // 初始化
    for i := 1; i <= amount; i++ {
        dp[i] = math.MaxInt32
    }

    for i := 0; i < len(coins); i++ {
        for j := 0; j <= amount; j++ {
            dp[j] = dp[j]
            if j >= coins[i] {
                dp[j] = min(dp[j], dp[j-coins[i]]+1)
            }
        }
    }
    if dp[amount] == math.MaxInt32 {
        return -1
    }

    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
