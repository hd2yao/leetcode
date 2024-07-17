package day30

func hasPathSum(root *TreeNode, targetSum int) bool {
    return pathSum(root, targetSum)
}

func pathSum(treeNode *TreeNode, sum int) bool {
    if treeNode == nil {
        return false
    }
    sum -= treeNode.Val
    if treeNode.Left == nil && treeNode.Right == nil {
        if sum == 0 {
            return true
        }
        return false
    }

    if treeNode.Left != nil {
        if pathSum(treeNode.Left, sum) {
            return true
        }
    }

    if treeNode.Right != nil {
        if pathSum(treeNode.Right, sum) {
            return true
        }
    }

    return false
}
