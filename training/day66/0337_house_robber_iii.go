package day66

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func robIII(root *TreeNode) int {
	dp := postorder(root)
	return max(dp[0], dp[1])
}

func postorder(treeNode *TreeNode) [2]int {
	if treeNode == nil {
		return [2]int{0, 0}
	}

	left := postorder(treeNode.Left)
	right := postorder(treeNode.Right)

	notRob := max(left[0], left[1]) + max(right[0], right[1])
	rob := treeNode.Val + left[0] + right[0]

	return [2]int{notRob, rob}
}
