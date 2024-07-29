package day42

func findSubsequences(nums []int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    var backtracking func(nums []int, startIndex int)
    backtracking = func(nums []int, startIndex int) {
        if len(path) > 1 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            result = append(result, tmp)
        }

        used := make(map[int]bool)

        for i := startIndex; i < len(nums); i++ {
            //if i != startIndex && nums[i] == nums[i-1] {
            //	continue
            //}
            if used[nums[i]] {
                continue
            }

            if len(path) == 0 || nums[i] >= path[len(path)-1] {
                used[nums[i]] = true
                path = append(path, nums[i])
                backtracking(nums, i+1)
                path = path[:len(path)-1]
            }
        }
    }

    backtracking(nums, 0)
    return result
}
