package day60

// 二维 dp
func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }

    if abs(target) > sum || (sum+target)%2 == 1 {
        return 0
    }

    target = (sum + target) / 2
    dp := make([][]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, target+1)
    }

    dp[0][0] = 1
    for i := 0; i < len(nums); i++ {
        for j := 0; j <= target; j++ {
            if i == 0 {
                if nums[i] == j {
                    dp[i][j]++
                }
            } else {
                if j >= nums[i] {
                    dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
                } else {
                    dp[i][j] = dp[i-1][j]
                }
            }
        }
    }

    return dp[len(nums)-2][target]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

// 一维 dp
func findTargetSumWaysOne(nums []int, target int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }

    if abs(target) > sum || (sum+target)%2 == 1 {
        return 0
    }

    target = (sum + target) / 2
    dp := make([]int, target+1)
    dp[0] = 1

    for i := 0; i < len(nums); i++ {
        for j := target; j >= nums[i]; j-- {
            dp[j] += dp[j-nums[i]]
        }
    }

    return dp[target]
}
