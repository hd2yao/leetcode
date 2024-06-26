package day9

import "sort"

func threeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)

    for i := 0; i < len(nums)-2; i++ {
        // 排序后首个元素大于 0，直接退出
        if nums[i] > 0 {
            break
        }
        // i 去重
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for left < right {
            if nums[i]+nums[left]+nums[right] == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                // 下面不能省略
                // left 去重
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // right 去重
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if nums[i]+nums[left]+nums[right] > 0 {
                // 剪枝可以省略
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                right--
            } else {
                // 剪枝可以省略
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                left++
            }
        }
    }
    return result
}
