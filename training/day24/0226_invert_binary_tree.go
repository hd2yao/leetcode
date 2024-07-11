package day24

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func invertTree(root *TreeNode) *TreeNode {
    invert(root)
    return root
}

func invert(treeNode *TreeNode) {
    if treeNode == nil {
        return
    }
    treeNode.Left, treeNode.Right = treeNode.Right, treeNode.Left
    invert(treeNode.Left)
    invert(treeNode.Right)
}
