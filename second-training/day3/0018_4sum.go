package day3

import "sort"

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result := make([][]int, 0)
    for i := 0; i < len(nums)-3; i++ {
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        for j := i + 1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j-1] == nums[j] {
                continue
            }
            left := j + 1
            right := len(nums) - 1
            for left < right {
                if nums[i]+nums[j]+nums[left]+nums[right] == target {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    for left < right && nums[left] == nums[left+1] {
                        left++
                    }
                    for left < right && nums[right] == nums[right-1] {
                        right--
                    }
                    left++
                    right--
                } else if nums[i]+nums[j]+nums[left]+nums[right] > target {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return result
}
