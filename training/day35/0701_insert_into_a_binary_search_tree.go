package day35

// 递归法
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }

    if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}

// 迭代法
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }

    node := root

    for node != nil {
        if node.Val < val {
            if node.Right != nil {
                node = node.Right
            } else {
                node.Right = &TreeNode{Val: val}
                break
            }
        } else {
            if node.Left != nil {
                node = node.Left
            } else {
                node.Left = &TreeNode{Val: val}
                break
            }
        }

    }

    return root
}
