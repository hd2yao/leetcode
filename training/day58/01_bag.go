package day58

// 0-1 背包：二维 dp

func bag(weight, value []int, bagWeight int) int {
    dp := make([][]int, len(weight))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, bagWeight+1)
    }

    for i := 0; i < len(weight); i++ {
        for j := 0; j <= bagWeight; j++ {
            // 初始化第一个物品
            if i == 0 {
                if j >= weight[i] {
                    dp[i][j] = value[i]
                }
            } else {
                if j >= weight[i] {
                    dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
                } else {
                    dp[i][j] = dp[i-1][j]
                }
            }
        }
    }

    return dp[len(weight)-1][bagWeight]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
