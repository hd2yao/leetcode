package day36

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

// 递归法
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root.Val < low {
        // 只考虑右子树
        if root.Right != nil {
            return trimBST(root.Right, low, high)
        }
        // 叶子节点 和 左子树 不考虑
        //if root.Left == nil && root.Right == nil {
        //    return nil
        //}
        return nil
    }
    if root.Val > high {
        if root.Left != nil {
            return trimBST(root.Left, low, high)
        }
        return nil
    }
    if root.Left != nil {
        root.Left = trimBST(root.Left, low, high)
    }
    if root.Right != nil {
        root.Right = trimBST(root.Right, low, high)
    }
    return root
}

// 简化代码
func trimBS2(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val < low {
        // 只考虑右子树
        return trimBST(root.Right, low, high)
    }
    if root.Val > high {
        return trimBST(root.Left, low, high)
    }
    root.Left = trimBST(root.Left, low, high)
    root.Right = trimBST(root.Right, low, high)
    return root
}
