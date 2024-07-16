package day29

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 0
    }

    leftValue := sumOfLeftLeaves(root.Left)
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        leftValue = root.Left.Val
    }
    rightValue := sumOfLeftLeaves(root.Right)

    sum := leftValue + rightValue
    return sum
}
