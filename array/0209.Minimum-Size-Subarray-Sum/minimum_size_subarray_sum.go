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
