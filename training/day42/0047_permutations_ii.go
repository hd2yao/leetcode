package day42

import "sort"

func permuteUnique(nums []int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(nums))

    sort.Ints(nums)

    var backtracking func(nums []int, used []bool)
    backtracking = func(nums []int, used []bool) {
        if len(path) == len(nums) {
            tmp := make([]int, len(path))
            copy(tmp, path)
            result = append(result, tmp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if i != 0 && nums[i] == nums[i-1] && used[i-1] == false {
                continue
            }
            if !used[i] {
                path = append(path, nums[i])
                used[i] = true
                backtracking(nums, used)
                used[i] = false
                path = path[:len(path)-1]
            }

        }
    }

    backtracking(nums, used)
    return result
}
