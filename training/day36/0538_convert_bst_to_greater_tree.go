package day36

func convertBST(root *TreeNode) *TreeNode {
    count := 0
    reverseInorder(root, &count)
    return root
}

func reverseInorder(root *TreeNode, count *int) {
    if root == nil {
        return
    }
    reverseInorder(root.Right, count)
    root.Val += *count
    *count = root.Val
    reverseInorder(root.Left, count)
}
