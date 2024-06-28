package main

func strStr(haystack string, needle string) int {
    return 0
}

func getNext(nextArray []int, needle string) {
    nextArray[0] = 0
    for i := 0; i < len(needle); i++ {
        if i == 0 {
            continue
        }
        subStrMap := make(map[string]int, 0)
        for j := 0; j < i; j++ {
            subStrMap[needle[:j+1]] = 1
        }
        for j := i; j > 0; j-- {
            if _, ok := subStrMap[needle[j:i+1]]; ok {
                nextArray[i] = max(nextArray[i], len(needle[1:j]))
            }
        }
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
