package day37

func combine(n int, k int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    var backtracking func(n, k, startIndex int)
    backtracking = func(n, k, startIndex int) {
        // 终止条件
        if len(path) == k {
            tmp := make([]int, k)
            copy(tmp, path)
            result = append(result, tmp)
            return
        }

        // 从 start 开始，不往回走
        for i := startIndex; i <= n; i++ {
            path = append(path, i)
            backtracking(n, k, i+1)
            path = path[:len(path)-1]
        }

    }

    backtracking(n, k, 1)
    return result
}
