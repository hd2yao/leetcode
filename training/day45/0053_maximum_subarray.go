package day45

func maxSubArray(nums []int) int {
    sum := 0
    max := nums[0]
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if max < sum {
            max = sum
        }
        if sum < 0 {
            sum = 0
            continue
        }
    }
    return max
}
