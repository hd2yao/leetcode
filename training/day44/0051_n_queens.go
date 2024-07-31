package day44

func solveNQueens(n int) [][]string {
    result := make([][]string, 0)
    path := make([]string, 0)

    // 记录每一个皇后的位置
    mark := make([][]int, n)
    for i := 0; i < len(mark); i++ {
        mark[i] = make([]int, n)
    }

    var backtracking func(n, row int)
    backtracking = func(n, row int) {
        if len(path) == n {
            tmp := make([]string, n)
            copy(tmp, path)
            result = append(result, tmp)
            return
        }

        for col := 0; col < n; col++ {
            // 判断当前行的皇后位置
            if !isValid(n, row, col, mark) {
                continue
            }

            mark[row][col] = 1

            // 构造当前行的皇后
            rowQueen := ""
            for j := 0; j < n; j++ {
                if j == col {
                    rowQueen += "Q"
                    continue
                }
                rowQueen += "."
            }
            path = append(path, rowQueen)

            // 下一行
            backtracking(n, row+1)

            path = path[:len(path)-1]
            mark[row][col] = 0
        }
    }

    backtracking(n, 0)
    return result
}

func isValid(n, row, col int, mark [][]int) bool {
    for i := 0; i < row; i++ {
        if mark[i][col] == 1 {
            return false
        }
    }

    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if mark[i][j] == 1 {
            return false
        }
    }

    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if mark[i][j] == 1 {
            return false
        }
    }

    return true
}
