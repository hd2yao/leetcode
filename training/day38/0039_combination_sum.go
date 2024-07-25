package day38

func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    combinationNums := make([]int, 0)

    var backtracking func(candidates []int, sum, target int)
    backtracking = func(candidates []int, sum, target int) {
        if sum == target {
            tmp := make([]int, len(combinationNums))
            copy(tmp, combinationNums)
            result = append(result, tmp)
            return
        }
        if sum > target {
            return
        }

        for i := 0; i < len(candidates); i++ {
            combinationNums = append(combinationNums, candidates[i])
            backtracking(candidates[i:], sum+candidates[i], target)
            combinationNums = combinationNums[:len(combinationNums)-1]
        }
    }

    backtracking(candidates, 0, target)
    return result
}
