package day31

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
    return traversal(inorder, postorder)
}

func traversal(inorder []int, postorder []int) *TreeNode {
    // 第一步：如果数组大小为零，说明是空节点
    if len(postorder) == 0 {
        return nil
    }

    // 第二步：如果不为空，后序数组的最后一个元素为根节点
    root := &TreeNode{Val: postorder[len(postorder)-1]}
    if len(postorder) == 1 {
        return root
    }

    // 第三步：找到后序数组的最后一个元素在中序数组中的位置，作为切割点
    delimiterIndex := 0
    for i, v := range inorder {
        if v == root.Val {
            delimiterIndex = i
            break
        }
    }

    // 第四步：切割中序数组，得到 中序左数组 和 中序右数组
    inorderLeft := inorder[:delimiterIndex]
    inorderRight := inorder[delimiterIndex+1:]

    // 第五步：切割后序数组，得到 后序左数组 和 后序右数组
    postorderLeft := postorder[:len(inorderLeft)]
    postorderRight := postorder[len(inorderLeft) : len(postorder)-1]

    // 第六步：
    root.Left = traversal(inorderLeft, postorderLeft)
    root.Right = traversal(inorderRight, postorderRight)

    return root
}
