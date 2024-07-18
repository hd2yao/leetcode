package day31

import "sort"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 1 找到最大值点
	index := maxIndex(nums)
	root := &TreeNode{Val: nums[index]}
	if len(nums) == 1 {
		return root
	}

	// 2 切割 nums 得到 左右子树数组
	numsLeft := nums[:index]
	numsRight := nums[index+1:]

	// 3 构造左右子树
	root.Left = constructMaximumBinaryTree(numsLeft)
	root.Right = constructMaximumBinaryTree(numsRight)

	return root
}

func maxIndex(nums []int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == numsCopy[len(numsCopy)-1] {
			index = i
			break
		}
	}
	return index
}

func maxIndex2(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}
