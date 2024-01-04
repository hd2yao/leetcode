package _018_4Sum

import "sort"

// nums[0] >= target 不可以这样去剪枝操作
// 例如 [-4, -1, 0, 0],  target = -5
// 这是因为 nums[0] >= target 默认 target 为正数，而负数是越加越小的

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return nil
    }
    sort.Ints(nums)
    result := [][]int{}
    for k := 0; k < len(nums)-3; k++ {
        // 一级剪枝
        if nums[k] > target && nums[k] >= 0 {
            break
        }
        // 一级去重
        if k > 0 && nums[k] == nums[k-1] {
            continue
        }
        for i := k + 1; i < len(nums)-2; i++ {
            // 二级剪枝
            if nums[k]+nums[i] > target && nums[i] >= 0 {
                // nums[k]+nums[i] > target && nums[k]+nums[i] >= 0 可以优化成上面的代码
                // 因为只要 nums[k]+nums[i] > target 同时 nums[i] 之后的数都是正数，四数之和就一定大于 target
                break
            }
            // 二级去重
            if i > k+1 && nums[i] == nums[i-1] {
                continue
            }
            left, right := i+1, len(nums)-1
            for left < right {
                if nums[k]+nums[i]+nums[left]+nums[right] == target {
                    result = append(result, []int{nums[k], nums[i], nums[left], nums[right]})
                    for left < right && nums[left] == nums[left+1] {
                        left++
                    }
                    for left < right && nums[right] == nums[right-1] {
                        right--
                    }
                    left++
                    right--
                } else if nums[k]+nums[i]+nums[left]+nums[right] < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return result
}
