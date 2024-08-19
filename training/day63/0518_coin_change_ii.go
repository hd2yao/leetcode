package day63

// 完全背包 二维 dp
func change(amount int, coins []int) int {
    dp := make([][]int, len(coins)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, amount+1)
    }
    dp[0][0] = 1

    for i := 1; i < len(coins)+1; i++ {
        for j := 0; j <= amount; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= coins[i-1] {
                dp[i][j] += dp[i][j-coins[i-1]]
            }
        }
    }
    return dp[len(coins)][amount]
}

// 完全背包 一维 dp
func changeOneDimension(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1
    for i := 0; i < len(coins); i++ {
        for j := coins[i]; j <= amount; j++ {
            dp[j] += dp[j-coins[i]]
        }
    }

    return dp[amount]
}
