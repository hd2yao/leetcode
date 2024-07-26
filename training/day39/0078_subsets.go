package day39

func subsets(nums []int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    result = append(result, []int{})

    var backtracking func(nums []int)
    backtracking = func(nums []int) {
        if nums == nil {
            return
        }

        for i := 0; i < len(nums); i++ {
            path = append(path, nums[i])

            tmp := make([]int, len(path))
            copy(tmp, path)
            result = append(result, tmp)

            backtracking(nums[i+1:])
            path = path[:len(path)-1]
        }
    }

    backtracking(nums)
    return result
}
