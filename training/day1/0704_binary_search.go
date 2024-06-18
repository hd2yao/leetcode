package day1

// 左闭右开区间
func searchNoRight(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        middle := (left + right) / 2
        if nums[middle] < target {
            left = middle + 1
        } else if nums[middle] > target {
            right = middle
        } else {
            return middle
        }
    }

    return -1
}

// 左闭右闭区间
func searchRight(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        middle := (left + right) / 2
        if nums[middle] < target {
            left = middle + 1
        } else if nums[middle] > target {
            right = middle
        } else {
            return middle
        }
    }

    return -1
}
