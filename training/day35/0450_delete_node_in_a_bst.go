package day35

func deleteNode(root *TreeNode, key int) *TreeNode {
    // 没有找到
    if root == nil {
        return nil
    }

    // 找到要删除节点的四种情况
    // 返回值为删除后的新节点
    if root.Val == key {
        // 叶子节点
        // 直接删除，所以返回 nil
        if root.Left == nil && root.Right == nil {
            return nil
        }
        // 左不空，右空
        // 将节点的左子树直接接入父节点的左子树中
        if root.Left != nil && root.Right == nil {
            return root.Left
        }
        // 左空，右不空
        // 将节点的右子树直接接入父节点的左子树中
        if root.Left == nil && root.Right != nil {
            return root.Right
        }
        // 左右都不空
        if root.Left != nil && root.Right != nil {
            // 先找到右子树的最左边的叶子节点
            cur := root.Right
            for cur.Left != nil {
                cur = cur.Left
            }
            cur.Left = root.Left
            return root.Right
        }
    }

    if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    }
    if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    }
    return root
}
