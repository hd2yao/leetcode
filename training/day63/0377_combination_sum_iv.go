package day63

// 二维 dp

func combinationSum4(nums []int, target int) int {
    dp := make([][]int, len(nums)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1

    for j := 0; j <= target; j++ {
        for i := 1; i <= len(nums); i++ {
            dp[i][j] = dp[i-1][j]
            if j >= nums[i-1] {
                dp[i][j] += dp[len(nums)][j-nums[i-1]]
            }
        }
    }
    return dp[len(nums)][target]
}
