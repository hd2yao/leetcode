package _059_Spiral_Matrix_II

func generateMatrix(n int) [][]int {
    // 矩阵初始化
    result := make([][]int, n)
    for i, _ := range result {
        result[i] = make([]int, n)
    }

    startX, startY := 0, 0
    offset := 1
    count := 1
    for circle := 0; circle < n/2; circle++ { // 循环的圈数
        // 左闭右开
        // 上行从左至右
        for j := startY; j < n-offset; j++ {
            result[startY][j] = count
            count++
        }
        // 右列从上至下
        for i := startX; i < n-offset; i++ {
            result[i][n-offset] = count
            count++
        }
        // 下行从右至左
        for j := n - offset; j > startY; j-- {
            result[n-offset][j] = count
            count++
        }
        // 左列从下至上
        for i := n - offset; i > startX; i-- {
            result[i][startX] = count
            count++
        }
        startX++
        startY++
        offset++
    }
    if n%2 == 1 {
        result[n/2][n/2] = count
    }
    return result
}
