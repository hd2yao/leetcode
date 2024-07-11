package day24

func isSymmetricTest(root *TreeNode) bool {
	node := *root
	invert1(&node)

	nodeVal := make([]int, 0)
	inorder(&node, &nodeVal)

	//nodeLevel := []*TreeNode{root}
	//for len(nodeLevel) > 0 {
	//	sizeLevel := len(nodeLevel)
	//	for i := 0; i < sizeLevel; i++ {
	//		if nodeLevel[0] != nil {
	//			nodeVal = append(nodeVal, nodeLevel[0].Val)
	//		} else {
	//			nodeVal = append(nodeVal, 101)
	//			continue
	//		}
	//		if nodeLevel[0].Left != nil {
	//			nodeLevel = append(nodeLevel, nodeLevel[0].Left)
	//		} else {
	//			nodeLevel = append(nodeLevel, nil)
	//		}
	//		if nodeLevel[0].Right != nil {
	//			nodeLevel = append(nodeLevel, nodeLevel[0].Right)
	//		} else {
	//			nodeLevel = append(nodeLevel, nil)
	//		}
	//		nodeLevel = nodeLevel[1:]
	//	}
	//}
	//
	rootVal := make([]int, 0)
	inorder(root, &rootVal)
	//rootLevel := []*TreeNode{root}
	//for len(rootLevel) > 0 {
	//	sizeLevel := len(rootLevel)
	//	for i := 0; i < sizeLevel; i++ {
	//		if rootLevel[0] != nil {
	//			rootVal = append(rootVal, rootLevel[0].Val)
	//		} else {
	//			rootVal = append(rootVal, 101)
	//			continue
	//		}
	//
	//		if rootLevel[0].Left != nil {
	//			rootLevel = append(rootLevel, rootLevel[0].Left)
	//		} else {
	//			rootLevel = append(rootLevel, nil)
	//		}
	//		if rootLevel[0].Right != nil {
	//			rootLevel = append(rootLevel, rootLevel[0].Right)
	//		} else {
	//			rootLevel = append(rootLevel, nil)
	//		}
	//		rootLevel = rootLevel[1:]
	//	}
	//}

	for i := 0; i < len(nodeVal); i++ {
		if nodeVal[i] != rootVal[i] {
			return false
		}
	}
	return true
}

func invert1(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}
	treeNode.Left, treeNode.Right = treeNode.Right, treeNode.Left
	invert(treeNode.Left)
	invert(treeNode.Right)
}

func inorder(treeNode *TreeNode, val *[]int) {
	if treeNode == nil {
		return
	}
	inorder(treeNode.Left, val)
	*val = append(*val, treeNode.Val)
	inorder(treeNode.Right, val)
}
