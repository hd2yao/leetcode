package day22

// 递归

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	postorder(root, &result)
	return result
}

func postorder(treeNode *TreeNode, val *[]int) {
	if treeNode == nil {
		return
	}
	postorder(treeNode.Left, val)
	postorder(treeNode.Right, val)
	*val = append(*val, treeNode.Val)
}
