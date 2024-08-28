package day72

func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }

    longest := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        if dp[i] > longest {
            longest = dp[i]
        }
    }
    return longest
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
