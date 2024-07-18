package day31

import "math"

// 使用中序遍历

func isValidBST(root *TreeNode) bool {
    inorderNums := make([]int, 0)
    inorder(root, &inorderNums)
    min := inorderNums[0]
    for i := 1; i < len(inorderNums); i++ {
        if inorderNums[i] <= min {
            return false
        }
        min = inorderNums[i]
    }
    return true
}

func inorder(treeNode *TreeNode, nums *[]int) {
    if treeNode == nil {
        return
    }
    inorder(treeNode.Left, nums)
    *nums = append(*nums, treeNode.Val)
    inorder(treeNode.Right, nums)
}

// 维护范围来确保每个节点都满足 BST 的条件

func isValidBST2(root *TreeNode) bool {
    return check(root, math.MinInt64, math.MaxInt64)
}

func check(treeNode *TreeNode, min, max int64) bool {
    if treeNode == nil {
        return true
    }
    if min >= int64(treeNode.Val) || max <= int64(treeNode.Val) {
        return false
    }

    return check(treeNode.Left, min, int64(treeNode.Val)) && check(treeNode.Right, int64(treeNode.Val), max)
}
