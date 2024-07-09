package day22

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

// 递归

func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    preorder(root, &result)
    return result
}

func preorder(treeNode *TreeNode, val *[]int) {
    if treeNode == nil {
        return
    }
    *val = append(*val, treeNode.Val)
    preorder(treeNode.Left, val)
    preorder(treeNode.Right, val)
}
