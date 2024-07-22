package day35

func lowestCommonAncestorSearchTree(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    // 找到节点
    if root == p || root == q {
        return root
    }

    // 区间中
    if (root.Val < p.Val && root.Val > q.Val) || (root.Val < q.Val && root.Val > p.Val) {
        return root
    }

    if root.Val > q.Val && root.Val > p.Val {
        leftTree := lowestCommonAncestorSearchTree(root.Left, p, q)
        if leftTree != nil {
            return leftTree
        }
    }

    if root.Val < q.Val && root.Val < p.Val {
        rightTree := lowestCommonAncestorSearchTree(root.Right, p, q)
        if rightTree != nil {
            return rightTree
        }
    }

    return root
}

// 代码简化
func lowestCommonAncestorSearchTree2(root, p, q *TreeNode) *TreeNode {
    if root.Val > q.Val && root.Val > p.Val {
        return lowestCommonAncestorSearchTree(root.Left, p, q)
    } else if root.Val < q.Val && root.Val < p.Val {
        return lowestCommonAncestorSearchTree(root.Right, p, q)
    }
    return root
}

// 迭代法
func lowestCommonAncestorSearchTree3(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if root.Val > q.Val && root.Val > p.Val {
            root = root.Left
        } else if root.Val < q.Val && root.Val < p.Val {
            root = root.Right
        } else {
            return root
        }
    }
    return nil
}
