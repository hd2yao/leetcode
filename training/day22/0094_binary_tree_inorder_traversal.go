package day22

// 递归

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorder(treeNode *TreeNode, val *[]int) {
	if treeNode == nil {
		return
	}
	inorder(treeNode.Left, val)
	*val = append(*val, treeNode.Val)
	inorder(treeNode.Right, val)
}
