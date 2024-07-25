package day38

import "sort"

func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    combinationNums := make([]int, 0)

    sort.Ints(candidates)

    var backtracking func(candidates []int, sum, target int)
    backtracking = func(candidates []int, sum, target int) {
        if sum == target {
            tmp := make([]int, len(combinationNums))
            copy(tmp, combinationNums)
            result = append(result, tmp)
            return
        }
        //if sum > target {
        //	return
        //}

        for i := 0; i < len(candidates); i++ {
            if sum+candidates[i] > target {
                break
            }
            combinationNums = append(combinationNums, candidates[i])
            backtracking(candidates[i:], sum+candidates[i], target)
            combinationNums = combinationNums[:len(combinationNums)-1]
        }
    }

    backtracking(candidates, 0, target)
    return result
}
