package day32

func findMode(root *TreeNode) []int {
    mode := make([]int, 0)
    nums := make([]int, 0)
    inorderNums(root, &nums)
    if len(nums) == 1 {
        return nums
    }
    maxCount := 1
    mode = append(mode, nums[0])
    for i := 1; i < len(nums); i++ {
        count := 1
        flag := false
        for i < len(nums) && nums[i] == nums[i-1] {
            flag = true
            count++
            i++
        }
        if maxCount < count {
            maxCount = count
            if flag {
                mode = []int{nums[i-1]}
            } else {
                mode = []int{nums[i]}
            }
        } else if maxCount == count {
            if flag {
                mode = append(mode, nums[i-1])
            } else {
                mode = append(mode, nums[i])
            }
        }
    }
    return mode
}

func inorderNums(treeNode *TreeNode, nums *[]int) {
    if treeNode == nil {
        return
    }
    inorderNums(treeNode.Left, nums)
    *nums = append(*nums, treeNode.Val)
    inorderNums(treeNode.Right, nums)
}

func modeCount(nums []int) []int {
    if len(nums) == 0 {
        return nil // 返回 nil 表示数组为空
    }

    var modes []int
    currentCount := 1
    maxCount := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            currentCount++
        } else {
            if currentCount > maxCount {
                maxCount = currentCount
                modes = []int{nums[i-1]}
            } else if currentCount == maxCount {
                modes = append(modes, nums[i-1])
            }
            currentCount = 1
        }
    }

    // 最后再检查一次 currentCount 是否大于 maxCount
    // 因为如果最后一个元素与前一个元素相同时，按照上面的逻辑会之间进入下一层循环，而此时已经遍历到最后一个元素会直接退出
    // 因此需要重新判断一下
    if currentCount > maxCount {
        modes = []int{nums[len(nums)-1]}
    } else if currentCount == maxCount {
        modes = append(modes, nums[len(nums)-1])
    }

    return modes
}

// 中序遍历过程中记录前一个节点直接比较

func findMode2(root *TreeNode) []int {
    mode := make([]int, 0)
    maxCount := 1
    count := 1
    var prev *TreeNode
    var travel func(node *TreeNode)
    travel = func(node *TreeNode) {
        if node == nil {
            return
        }
        travel(node.Left)
        if prev != nil && node.Val == prev.Val {
            count++
        } else {
            count = 1
        }
        if count > maxCount {
            mode = []int{node.Val}
            maxCount = count
        } else if count == maxCount {
            mode = append(mode, node.Val)
        }
        prev = node
        travel(node.Right)
    }

    travel(root)
    return mode
}
