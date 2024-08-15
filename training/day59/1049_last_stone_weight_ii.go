package day59

func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[target] - dp[target]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 二维 dp
func lastStoneWeightII2(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2

	dp := make([][]int, len(stones))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < len(stones); i++ {
		for j := 0; j <= target; j++ {
			if i == 0 {
				if j >= stones[i] {
					dp[i][j] = stones[i]
				}
			} else {
				if j >= stones[i] {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return sum - dp[len(stones)-1][target] - dp[len(stones)-1][target]
}
