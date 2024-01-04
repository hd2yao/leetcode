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

// 解法二 kSum
func fourSum1(nums []int, target int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)
	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSum(nums []int, left, right int, target int, k int, cur []int, res *[][]int) {
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}
	if k == 2 {
		// 2 sum
		twoSum(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))
				copy(next, cur)
				next = append(next, nums[i])
				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}

}

func twoSum(nums []int, left, right int, target int, cur []int, res *[][]int) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)
			// reset cur to previous state
			cur = cur[:len(cur)-2]
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}

// 解法三 map
func fourSum2(nums []int, target int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*3+uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*3+uniqNums[i] == target) && counter[uniqNums[j]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i]*2 == target) && counter[uniqNums[j]] > 1 && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[j]*2+uniqNums[i]+uniqNums[k] == target) && counter[uniqNums[j]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[k]*2+uniqNums[i]+uniqNums[j] == target) && counter[uniqNums[k]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return res
}
