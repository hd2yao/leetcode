package day2

func minSubArrayLen(target int, nums []int) int {
    // 分别为子数组起始下标、子数组终止下标、子数组总和、目标子数组的最小长度
    left, right, sum, length := 0, 0, 0, 0
    for left < len(nums) || right < len(nums) {
        if left == 0 && right == 0 {
            sum = nums[0]
        }
        if sum < target {
            right++
            if right == len(nums) {
                break
            }
            sum += nums[right]
        } else if sum >= target {
            curLen := right - left + 1
            if length == 0 {
                length = curLen
            } else if curLen < length {
                length = curLen
            }
            sum -= nums[left]
            left++
        }
    }
    return length
}
