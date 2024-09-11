package day1

// 控制滑动窗口右边界
func minSubArrayLenRight(target int, nums []int) int {
    right := 0
    minLength := len(nums) + 1
    total := nums[0]

    for i := 0; i < len(nums); i++ {
        for total < target {
            right++
            if right == len(nums) {
                goto result
            }
            total += nums[right]
        }
        if minLength > right-i+1 {
            minLength = right - i + 1
        }
        total -= nums[i]
    }
result:
    if minLength == len(nums)+1 {
        return 0
    }
    return minLength
}

// 控制滑动窗口左边界
func minSubArrayLenLeft(target int, nums []int) int {
    left := 0
    minLength := 0
    total := 0

    for i := 0; i < len(nums); i++ {
        total += nums[i]
        for total >= target {
            if minLength == 0 || (minLength != 0 && minLength > i-left+1) {
                minLength = i - left + 1
            }
            total -= nums[left]
            left++
        }
    }

    return minLength
}
