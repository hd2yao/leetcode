package day1

/**

  本题中，当滑动窗口的右边界到数组末尾时就结束了
  因此，当我们控制右边界，用 for 来控制左边界时，会需要跳出；
  而用 for 来控制右边界时不需要

*/

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
