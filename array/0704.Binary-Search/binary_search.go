package _704_Binary_Search

// 解法一：左闭右闭
func search(nums []int, target int) int {
    lens := len(nums)
    if target < nums[0] || target > nums[lens-1] || lens == 0 {
        return -1
    }
    left, right := 0, lens-1
    // [left, right] left=right 仍然是合法的区间
    for left <= right {
        middle := (left + right) / 2
        if nums[middle] == target {
            return middle
        } else if nums[middle] > target {
            right = middle - 1
        } else {
            left = middle + 1
        }
    }
    return -1
}

// 解法二：左闭右开
func searchNoRight(nums []int, target int) int {
    lens := len(nums)
    if target < nums[0] || target > nums[lens-1] || lens == 0 {
        return -1
    }
    left, right := 0, lens
    // [left, right) left=right 不再是合法的区间
    for left < right {
        middle := (left + right) / 2
        if nums[middle] == target {
            return middle
        } else if nums[middle] > target {
            right = middle
        } else {
            left = middle + 1
        }
    }
    return -1
}
