package day2

func sortedSquares(nums []int) []int {
    front, rear := 0, len(nums)-1
    result := make([]int, len(nums))
    index := len(nums) - 1
    for front <= rear {
        if abs(nums[front]) <= abs(nums[rear]) {
            // square(nums[front]) <= square(nums[rear])
            result[index] = square(nums[rear])
            rear--

        } else {
            result[index] = square(nums[front])
            front++
        }
        index--
    }
    return result
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
