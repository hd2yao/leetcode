package day52

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func minCameraCover(root *TreeNode) int {
    count := 0

    var postorder func(treeNode *TreeNode, count *int) (mark int)
    postorder = func(treeNode *TreeNode, count *int) (mark int) {
        if treeNode == nil {
            return 2
        }
        leftMark := postorder(treeNode.Left, count)
        rightMark := postorder(treeNode.Right, count)
        if leftMark == 2 && rightMark == 2 {
            return 0
        } else if leftMark == 0 || rightMark == 0 {
            *count++
            return 1
        }
        return 2
    }

    // 说明根节点也是无覆盖的
    if postorder(root, &count) == 0 {
        count++
    }
    return count
}
