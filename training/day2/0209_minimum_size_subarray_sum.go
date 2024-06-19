package day2

// 滑动窗口-双指针
func minSubArrayLen(target int, nums []int) int {
    // 分别为子数组起始下标、子数组终止下标、子数组总和、目标子数组的最小长度
    left, right, length := 0, 0, 0
    sum := nums[0]
    for left < len(nums) || right < len(nums) {
        if sum < target {
            right++
            if right == len(nums) {
                break
            }
            sum += nums[right]
        } else {
            curLen := right - left + 1
            if length == 0 || (length != 0 && curLen < length) {
                length = curLen
            }
            sum -= nums[left]
            left++
        }
    }
    return length
}

// 滑动窗口-优化 单指针
func minSubArrayLenOptimize(target int, nums []int) int {
    // left 为起始位置
    left, length := 0, 0
    sum := 0
    // 终止位置通过 for 循环来移动
    for right, num := range nums {
        sum += num
        for sum >= target {
            curLen := right - left + 1
            if length == 0 || (length != 0 && curLen < length) {
                length = curLen
            }
            sum -= nums[left]
            left++
        }

    }
    return length
}

// 滑动窗口-优化 单指针
func minSubArrayLenOptimizeLeft(target int, nums []int) int {
    // left 为起始位置
    right, length := 0, 0
    sum := nums[0]
    // 起始位置通过 for 循环来移动
    for left, num := range nums {
        for sum < target {
            right++
            if right < len(nums) {
                sum += nums[right]
            } else {
                goto out
            }
        }
        curLen := right - left + 1
        if length == 0 || (length != 0 && curLen < length) {
            length = curLen
        }
        sum -= num
    }
out:
    return length
}
