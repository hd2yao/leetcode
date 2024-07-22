package day35

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == p || root == q || root == nil {
        return root
    }

    leftTree := lowestCommonAncestor(root.Left, p, q)
    rightTree := lowestCommonAncestor(root.Right, p, q)

    if leftTree != nil && rightTree != nil {
        return root
    }
    if leftTree == nil {
        return rightTree
    }
    if rightTree == nil {
        return leftTree
    }
    return nil
}
