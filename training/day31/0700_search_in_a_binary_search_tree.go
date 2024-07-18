package day31

// 迭代法
func searchBST(root *TreeNode, val int) *TreeNode {
    for root != nil {
        if root.Val == val {
            return root
        } else if root.Val > val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return nil
}

// 递归法
func search(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    if root.Val > val {
        return search(root.Left, val)
    }
    return search(root.Right, val)
}
