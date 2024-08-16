package day60

func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 0
    for i := 0; i < len(strs); i++ {
        zero, one := 0, 0
        for _, c := range strs[i] {
            if c == '0' {
                zero++
            } else {
                one++
            }
        }
        for j := m; j >= zero; j-- {
            for k := n; k >= one; k-- {
                dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
            }
        }
    }

    return dp[m][n]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
