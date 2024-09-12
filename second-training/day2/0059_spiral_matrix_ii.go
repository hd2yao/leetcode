package day2

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	count := 1
	for loop := 0; loop < n/2; loop++ {
		// 上行：从左至右
		// j 从 left 开始，可能能清晰易懂
		for j := loop; j < n-1-loop; j++ {
			matrix[top][j] = count
			count++
		}
		// 右列：从上至下
		// i := top
		for i := loop; i < n-1-loop; i++ {
			matrix[i][right] = count
			count++
		}
		// 下行：从右至左
		// j > left
		for j := n - 1 - loop; j > loop; j-- {
			matrix[bottom][j] = count
			count++
		}
		// 左列：从下至上
		// i > top
		for i := n - 1 - loop; i > loop; i-- {
			matrix[i][left] = count
			count++
		}
		left++
		right--
		top++
		bottom--
	}

	if n%2 == 1 {
		matrix[n/2][n/2] = n * n
	}

	return matrix
}
