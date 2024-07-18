package day31

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
