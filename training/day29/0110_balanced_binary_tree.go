package day29

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func isBalanced(root *TreeNode) bool {
    height := getHeight(root)
    if height == -1 {
        return false
    }
    return true
}

func getHeight(treeNode *TreeNode) int {
    if treeNode == nil {
        return 0
    }
    lh, rh := getHeight(treeNode.Left), getHeight(treeNode.Right)
    if lh == -1 || rh == -1 {
        return -1
    }

    if lh-rh > 1 || rh-lh > 1 {
        return -1
    }

    return max(lh, rh) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
