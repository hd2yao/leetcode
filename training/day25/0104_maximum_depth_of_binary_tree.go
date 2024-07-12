package day25

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

// 迭代法 -- 层序遍历

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    currentLevel := []*TreeNode{root}
    level := 0
    for len(currentLevel) > 0 {
        sizeLevel := len(currentLevel)
        for i := 0; i < sizeLevel; i++ {
            if currentLevel[0].Left != nil {
                currentLevel = append(currentLevel, currentLevel[0].Left)
            }
            if currentLevel[0].Right != nil {
                currentLevel = append(currentLevel, currentLevel[0].Right)
            }
            currentLevel = currentLevel[1:]
        }
        level++
    }
    return level
}

// 递归法 -- 前序或后续遍历

func maxDepth1(root *TreeNode) int {
    return getDepth(root)
}

func getDepth(treeNode *TreeNode) int {
    if treeNode == nil {
        return 0
    }
    left := getDepth(treeNode.Left)
    right := getDepth(treeNode.Right)
    return max(left, right) + 1
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
