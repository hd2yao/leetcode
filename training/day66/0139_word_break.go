package day66

// 二维 dp
func wordBreak(s string, wordDict []string) bool {
    dp := make([][]bool, len(wordDict)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, len(s)+1)
        dp[i][0] = true
    }

    for j := 1; j <= len(s); j++ {
        for i := 1; i <= len(wordDict); i++ {
            dp[i][j] = dp[i-1][j]
            if j >= len(wordDict[i-1]) {
                dp[i][j] = dp[i][j] || (s[j-len(wordDict[i-1]):j] == wordDict[i-1] && dp[len(wordDict)][j-len(wordDict[i-1])])
            }
        }
    }

    return dp[len(wordDict)][len(s)]
}

// 一维 dp
func wordBreakOneDimension(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for j := 1; j <= len(s); j++ {
        for i := 0; i < len(wordDict); i++ {
            if j >= len(wordDict[i]) && dp[j-len(wordDict[i])] && (s[j-len(wordDict[i]):j] == wordDict[i]) {
                dp[j] = true
                break
            }
        }
    }

    return dp[len(s)]
}
