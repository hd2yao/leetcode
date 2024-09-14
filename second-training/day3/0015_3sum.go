package day3

import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    if nums[0] > 0 || nums[len(nums)-1] < 0 {
        return nil
    }
    result := make([][]int, 0)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for left < right {
            if nums[i]+nums[left]+nums[right] == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[right-1] == nums[right] {
                    right--
                }
                for left < right && nums[left+1] == nums[left] {
                    left++
                }
                left++
                right--
            } else if nums[i]+nums[left]+nums[right] > 0 {
                right--
            } else {
                left++
            }
        }
    }
    return result
}
