package day30

func pathSumII(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	pathSumResult(root, targetSum, new([]int), &result)
	return result
}

func pathSumResult(treeNode *TreeNode, sum int, path *[]int, result *[][]int) {
	sum -= treeNode.Val
	*path = append(*path, treeNode.Val)
	if treeNode.Left == nil && treeNode.Right == nil && sum == 0 {
		pathCopy := make([]int, len(*path))
		copy(pathCopy, *path)
		*result = append(*result, pathCopy)
		return
	}

	if treeNode.Left != nil {
		pathSumResult(treeNode.Left, sum, path, result)
	}

	if treeNode.Right != nil {
		pathSumResult(treeNode.Right, sum, path, result)
	}
	*path = (*path)[:len(*path)-1]
}
