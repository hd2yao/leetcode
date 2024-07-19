package day32

import (
    "math"

    "github.com/hd2yao/leetcode/training/structures"
)

type TreeNode = structures.TreeNode

// 中序遍历得到数组

func getMinimumDifference(root *TreeNode) int {
    nums := make([]int, 0)
    inorder(root, &nums)
    min := math.MaxInt
    for i := 1; i < len(nums); i++ {
        if min > nums[i]-nums[i-1] {
            min = nums[i] - nums[i-1]
        }
    }
    return min
}

func inorder(treeNode *TreeNode, nums *[]int) {
    if treeNode == nil {
        return
    }
    inorder(treeNode.Left, nums)
    *nums = append(*nums, treeNode.Val)
    inorder(treeNode.Right, nums)
}
