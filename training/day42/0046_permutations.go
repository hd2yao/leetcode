package day42

func permute(nums []int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)
    used := make([]int, len(nums))

    var backtracking func(nums []int, used []int)
    backtracking = func(nums []int, used []int) {
        if len(path) == len(nums) {
            tmp := make([]int, len(path))
            copy(tmp, path)
            result = append(result, tmp)

            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] == 1 {
                continue
            }
            path = append(path, nums[i])
            used[i] = 1
            backtracking(nums, used)
            path = path[:len(path)-1]
            used[i] = 0
        }
    }

    backtracking(nums, used)
    return result
}
