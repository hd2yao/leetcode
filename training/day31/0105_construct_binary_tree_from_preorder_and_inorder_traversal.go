package day31

func buildTree2(preorder []int, inorder []int) *TreeNode {
    // 1. 判断是否为空节点
    if len(preorder) == 0 {
        return nil
    }

    // 2. 根据前序遍历构造根节点
    root := &TreeNode{Val: preorder[0]}
    if len(preorder) == 1 {
        return root
    }

    // 3. 找到根节点在中序遍历中的位置，作为切割点
    delimiterIndex := 0
    for i, v := range inorder {
        if v == root.Val {
            delimiterIndex = i
            break
        }
    }

    // 4. 切割中序遍历数组
    inorderLeft := inorder[:delimiterIndex]
    inorderRight := inorder[delimiterIndex+1:]

    // 5. 切割前序遍历数组
    preorderLeft := preorder[1 : 1+len(inorderLeft)]
    preorderRight := preorder[1+len(inorderLeft):]

    // 6. 递归构造左右子树
    root.Left = buildTree2(preorderLeft, inorderLeft)
    root.Right = buildTree2(preorderRight, inorderRight)

    return root
}
