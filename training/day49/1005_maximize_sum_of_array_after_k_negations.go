package day49

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
    sort.Ints(nums)

    for k > 0 {
        if nums[0] < 0 {
            nums[0] = -nums[0]
            k--
        } else if nums[0] > 0 {
            if k%2 == 1 {
                nums[0] = -nums[0]
            }
            break
        } else {
            break
        }
        sort.Ints(nums)
    }

    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }
    return sum
}
