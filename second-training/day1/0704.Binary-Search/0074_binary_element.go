package _704_Binary_Search

// 左闭右开
func search(nums []int, target int) int {
    // 可以不加，提前判断可以减少后面的过程
    if nums[0] > target || nums[len(nums)-1] < target {
        return -1
    }

    // right 取不到值，因此赋值为 len(nums)
    left, right := 0, len(nums)

    for left < right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return -1
}

// 左闭右闭
func search2(nums []int, target int) int {
    // 可以不加，提前判断可以减少后面的过程
    if nums[0] > target || nums[len(nums)-1] < target {
        return -1
    }

    // right 可以取到值，因此赋值为 len(nums) - 1
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
