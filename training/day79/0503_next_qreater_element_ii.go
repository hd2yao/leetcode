package day79

// 拼接
func nextGreaterElements(nums []int) []int {
    newNums := make([]int, 0)
    newNums = append(newNums, nums...)
    newNums = append(newNums, nums...)

    stack := make([]int, 0)
    answer := make([]int, len(newNums))
    for i := 0; i < len(nums); i++ {
        answer[i] = -1
    }

    for i := 0; i < len(newNums); i++ {
        for len(stack) != 0 && newNums[stack[len(stack)-1]] < newNums[i] {
            index := stack[len(stack)-1]
            answer[index] = newNums[i]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }

    return answer[:len(nums)]
}

// 模拟拼接过程
func nextGreaterElements2(nums []int) []int {
    stack := make([]int, 0)
    answer := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        answer[i] = -1
    }

    for i := 0; i < len(nums)*2; i++ {
        for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i%len(nums)] {
            index := stack[len(stack)-1]
            answer[index] = nums[i%len(nums)]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i%len(nums))
    }

    return answer
}
