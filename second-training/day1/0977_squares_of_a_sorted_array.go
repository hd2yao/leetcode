package day1

func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))

    left, right := 0, len(nums)-1
    for i := len(result) - 1; i >= 0; i-- {
        if square(nums[right]) > square(nums[left]) {
            result[i] = square(nums[right])
            right--
        } else {
            result[i] = square(nums[left])
            left++
        }
    }
    return result
}

func square(num int) int {
    return num * num
}
