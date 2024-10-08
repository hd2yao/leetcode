package day58

func canPartition(nums []int) bool {
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }
    if sum%2 == 1 {
        return false
    }

    sum /= 2
    dp := make([]int, sum+1)
    for i := 0; i < len(nums); i++ {
        for j := sum; j >= nums[i]; j-- {
            dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
        }
    }

    return dp[sum] == sum
}
