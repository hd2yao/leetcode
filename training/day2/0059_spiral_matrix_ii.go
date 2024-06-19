package day2

func generateMatrix(n int) [][]int {
    result := make([][]int, n)
    for index, _ := range result {
        result[index] = make([]int, n)
    }
    top, bottom := 0, n-1
    left, right := 0, n-1

    offset := 1
    count := 1
    for circle := 0; circle < (n / 2); circle++ {
        // 上行：从左至右
        for y := left; y < n-offset; y++ {
            result[top][y] = count
            count++
        }
        // 右列：从上至下
        for x := top; x < n-offset; x++ {
            result[x][right] = count
            count++
        }
        // 下行：从右至左
        for y := n - offset; y > left; y-- {
            result[bottom][y] = count
            count++
        }
        // 左列：从下至上
        for x := n - offset; x > top; x-- {
            result[x][left] = count
            count++
        }
        offset++
        top++
        bottom--
        left++
        right--
    }

    if n%2 == 1 {
        result[n/2][n/2] = count
    }
    return result
}
