package day29

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	preorderPath(root, "", &result)
	return result
}

func preorderPath(treeNode *TreeNode, path string, paths *[]string) {
	if treeNode.Left == nil && treeNode.Right == nil {
		path += strconv.Itoa(treeNode.Val)
		*paths = append(*paths, path)
		return
	}
	path += strconv.Itoa(treeNode.Val) + "->"
	if treeNode.Left != nil {
		preorderPath(treeNode.Left, path, paths)
	}

	if treeNode.Right != nil {
		preorderPath(treeNode.Right, path, paths)
	}
}
