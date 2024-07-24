package day37

// 未剪枝

func combinationSum3(k int, n int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    var backtracking func(k, n, startIndex int)
    backtracking = func(k, n, startIndex int) {
        if len(path) == k {
            sum := 0
            for i := 0; i < len(path); i++ {
                sum += path[i]
            }
            if sum == n {
                tmp := make([]int, k)
                copy(tmp, path)
                result = append(result, tmp)
            }
            return
        }

        for i := startIndex; i <= 9; i++ {
            path = append(path, i)
            backtracking(k, n, i+1)
            path = path[:len(path)-1]
        }
    }

    backtracking(k, n, 1)
    return result
}

// 剪枝操作

func combinationSum32(k int, n int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)

    var backtracking func(k, sum, n, startIndex int)
    backtracking = func(k, sum, n, startIndex int) {
        if len(path) == k {
            if sum == n {
                tmp := make([]int, k)
                copy(tmp, path)
                result = append(result, tmp)
            }
            return
        }

        for i := startIndex; i <= 9-(k-len(path))+1; i++ {
            if sum+i > n {
                break
            }
            path = append(path, i)
            backtracking(k, sum+i, n, i+1)
            path = path[:len(path)-1]
        }
    }

    backtracking(k, 0, n, 1)
    return result
}
