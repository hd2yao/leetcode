package day38

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, 0)

	sort.Ints(candidates)

	//var backtracking func(candidates []int, sum, target int)
	//backtracking = func(candidates []int, sum, target int) {
	//	if sum == target {
	//		tmp := make([]int, len(nums))
	//		copy(tmp, nums)
	//		result = append(result, tmp)
	//		return
	//	}
	//
	//	for i := 0; i < len(candidates); i++ {
	//		if i != 0 && candidates[i] == candidates[i-1] {
	//			continue
	//		}
	//		if sum+candidates[i] > target {
	//			break
	//		}
	//		nums = append(nums, candidates[i])
	//		backtracking(candidates[i+1:], sum+candidates[i], target)
	//		nums = nums[:len(nums)-1]
	//	}
	//}
	var backtracking func(candidates []int, startIndex, sum, target int)
	backtracking = func(candidates []int, startIndex, sum, target int) {
		if sum == target {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			if i != startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			if sum+candidates[i] > target {
				break
			}
			nums = append(nums, candidates[i])
			backtracking(candidates, i+1, sum+candidates[i], target)
			nums = nums[:len(nums)-1]
		}
	}

	backtracking(candidates, 0, 0, target)
	return result
}
