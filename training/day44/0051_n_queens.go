package day44

func solveNQueens(n int) [][]string {
    result := make([][]string, 0)
    path := make([]string, 0)

    // 记录每一个皇后的位置
    rowMark := make([]int, n)
    colMark := make([]int, n)

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

            // 当前列已有皇后
            if colMark[col] == 1 {
                continue
            }

            // 对角线上已有皇后
            // 只通过行列无法确定对角线
            if row != 0 && rowMark[row-1] == 1 && ((col != 0 && col != n-1 && (colMark[col-1] == 1 || colMark[col+1] == 1)) || (col == 0 && colMark[col+1] == 1) || (col == n-1 && colMark[col-1] == 1)) {
                continue
            }

            rowMark[row] = 1
            colMark[col] = 1

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
            rowMark[row] = 0
            colMark[col] = 0
        }
    }

    backtracking(n, 0)
    return result
}
