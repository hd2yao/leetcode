package day72

func findLengthOfLCIS(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = 1
    longest := 1
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            dp[i] = dp[i-1] + 1
        } else {
            dp[i] = 1
        }

        if longest < dp[i] {
            longest = dp[i]
        }
    }
    return longest
}
