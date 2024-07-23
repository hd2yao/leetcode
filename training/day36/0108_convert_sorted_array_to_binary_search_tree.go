package day36

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    delimiterIndex := len(nums) / 2
    root := &TreeNode{Val: nums[delimiterIndex]}
    root.Left = sortedArrayToBST(nums[:delimiterIndex])
    root.Right = sortedArrayToBST(nums[delimiterIndex+1:])
    return root
}
