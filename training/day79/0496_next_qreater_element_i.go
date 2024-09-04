package day79

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	answer := make(map[int]int, 0)
	stack := make([]int, 0)

	for i := 0; i < len(nums2); i++ {
		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			answer[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	stack = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		stack[i] = answer[nums1[i]]
		if stack[i] == 0 {
			stack[i] = -1
		}
	}
	return stack
}
