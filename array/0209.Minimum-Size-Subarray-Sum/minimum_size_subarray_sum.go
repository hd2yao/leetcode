package _209_Minimum_Size_Subarray_Sum

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
    left, right := 0, 0
    sum := nums[0]
    subArrayLens := 0
    for left <= right && right < len(nums) {
        //fmt.Println(left, right, sum, subArrayLens)
        if sum >= target {
            if subArrayLens == 0 || (subArrayLens != 0 && subArrayLens > right-left+1) {
                subArrayLens = right - left + 1
            }
            sum -= nums[left]
            left++
        } else {
            if right == len(nums)-1 {
                break
            }
            right++
            sum += nums[right]
        }

    }
    //fmt.Println("final", left, right, sum, subArrayLens)
    return subArrayLens
}

// 优化代码
func minSubArrayLenOptimize(target int, nums []int) int {
    left := 0
    sum := 0
    subArrayLens := 0
    for right, num := range nums {
        // 移动终止位置
        sum += num
        // sum >= target 时，找到一个解
        for sum >= target {
            // 这里其实就是 min()
            if subArrayLens == 0 || (subArrayLens != 0 && subArrayLens > right-left+1) {
                subArrayLens = right - left + 1
            }
            // 移动起始位置，去寻找最优解
            sum -= nums[left]
            left++
        }
    }
    //fmt.Println("final", left, sum, subArrayLens)
    return subArrayLens
}
