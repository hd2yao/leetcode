package _977_Squares_of_a_Sorted_Array

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	index := len(nums) - 1
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			res[index] = square(nums[left])
			break
		}
		if abs(nums[left]) < abs(nums[right]) {
			res[index] = square(nums[right])
			right--
			index--
		} else {
			res[index] = square(nums[left])
			left++
			index--
		}
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
