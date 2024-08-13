package day57

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			curMax := max(j*(i-j), j*dp[i-j])
			dp[i] = max(dp[i], curMax)
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
