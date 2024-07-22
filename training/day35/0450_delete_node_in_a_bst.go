package day35

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 没有找到
	if root == nil {
		return nil
	}

	if root.Val == key {
		// 叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 左不空，右空
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// 左空，右不空
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
