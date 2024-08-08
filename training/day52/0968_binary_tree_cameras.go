package day52

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func minCameraCover(root *TreeNode) int {
    count := 0

    var postorder func(treeNode *TreeNode, count *int) (mark int)
    postorder = func(treeNode *TreeNode, count *int) (mark int) {

        // 空节点，该节点有覆盖
        if treeNode == nil {
            return 2
        }

        leftMark := postorder(treeNode.Left, count)
        rightMark := postorder(treeNode.Right, count)

        // 情况1
        // 左右节点都有覆盖
        if leftMark == 2 && rightMark == 2 {
            return 0
        }

        // 情况2
        // left == 0 && right == 0 左右节点无覆盖
        // left == 1 && right == 0 左节点有摄像头，右节点无覆盖
        // left == 0 && right == 1 左节点有无覆盖，右节点摄像头
        // left == 0 && right == 2 左节点无覆盖，右节点覆盖
        // left == 2 && right == 0 左节点覆盖，右节点无覆盖
        if leftMark == 0 || rightMark == 0 {
            *count++
            return 1
        }

        // 情况3
        // left == 1 && right == 2 左节点有摄像头，右节点有覆盖
        // left == 2 && right == 1 左节点有覆盖，右节点有摄像头
        // left == 1 && right == 1 左右节点都有摄像头
        // 其他情况前段代码均已覆盖
        return 2
    }

    // 说明根节点也是无覆盖的
    if postorder(root, &count) == 0 {
        count++
    }
    return count
}
