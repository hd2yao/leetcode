package day44

func solveSudoku(board [][]byte) {
    var backtracking func(board [][]byte) bool
    backtracking = func(board [][]byte) bool {
        for i := 0; i < 9; i++ { // 遍历行
            for j := 0; j < 9; j++ { // 遍历列
                if board[i][j] != '.' {
                    continue
                }

                // 空位从 1 开始尝试
                for k := '1'; k <= '9'; k++ {
                    // 判断当前位置是否合适放 k，不合适直接下一个数
                    if isvalid(i, j, byte(k), board) {
                        // 放入 k
                        board[i][j] = byte(k)
                        // 然后递归进行下一个数的放置
                        // 因为只有一个解，所以只要找到一组就返回
                        if backtracking(board) {
                            return true
                        }
                        // 回溯
                        board[i][j] = '.'
                    }
                }

                // 说明九个数都尝试了，没有可以的数
                return false
            }
        }

        return true
    }

    backtracking(board)
}

func isvalid(row, col int, k byte, board [][]byte) bool {
    // 同一行重复
    for i := 0; i < 9; i++ {
        if board[row][i] == k {
            return false
        }
    }

    // 同一列重复
    for i := 0; i < 9; i++ {
        if board[i][col] == k {
            return false
        }
    }

    // 同一个九宫格重复
    startRow := (row / 3) * 3
    startCol := (col / 3) * 3
    for i := startRow; i < startRow+3; i++ {
        for j := startCol; j < startCol+3; j++ {
            if board[i][j] == k {
                return false
            }
        }
    }

    return true
}
