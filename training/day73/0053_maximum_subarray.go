package day73

func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    result := dp[0]
    for i := 1; i < len(dp); i++ {
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        if result < dp[i] {
            result = dp[i]
        }
    }
    return result
}
