package day66

func robII(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    money1 := rob(nums[1:])
    money2 := rob(nums[:len(nums)-1])

    return max(money1, money2)
}
