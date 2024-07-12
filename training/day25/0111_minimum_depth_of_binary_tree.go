package day25

// 递归法

func minDepth(root *TreeNode) int {
    return getMinDepth(root)
}

func getMinDepth(treeNode *TreeNode) int {
    if treeNode == nil {
        return 0
    }
    if treeNode.Left == nil && treeNode.Right != nil {
        return 1 + getMinDepth(treeNode.Right)
    }
    if treeNode.Right == nil && treeNode.Left != nil {
        return 1 + getMinDepth(treeNode.Left)
    }
    return 1 + min(getMinDepth(treeNode.Left), getMinDepth(treeNode.Right))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// 迭代法

func minDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    depth := 0
    for len(queue) > 0 {
        size := len(queue)
        depth++
        for i := 0; i < size; i++ {
            if queue[0].Left == nil && queue[0].Right == nil {
                return depth
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
            }
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
            }
            queue = queue[1:]
        }
    }
    return depth
}
