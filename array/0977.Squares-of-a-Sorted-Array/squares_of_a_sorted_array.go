package _977_Squares_of_a_Sorted_Array

func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    left, right := 0, len(nums)
    for left < right {
        right--
        if left == right {
            res[left] = square(nums[left])
            break
        }
        if abs(nums[left]) < abs(nums[right]) {
            res[right] = square(nums[right])
            continue
        }
        res[left], res[right] = nums[right], square(nums[left])

    }
    return res
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func square(num int) int {
    return num * num
}

func returnResult(a, b int) int {
    if abs(a) < abs(b) {
        return square(a)
    }
    return square(b)
}
