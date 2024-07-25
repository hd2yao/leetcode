package day38

func partition(s string) [][]string {
    result := make([][]string, 0)
    path := make([]string, 0)

    var backtracking func(s string)
    backtracking = func(s string) {
        if s == "" {
            tmp := make([]string, len(path))
            copy(tmp, path)
            result = append(result, tmp)
            return
        }

        for i := 0; i < len(s); i++ {
            if !isPalindrome(s[:i+1]) {
                continue
            }
            path = append(path, s[:i+1])
            backtracking(s[i+1:])
            path = path[:len(path)-1]
        }
    }

    backtracking(s)
    return result
}

func isPalindrome(s string) bool {
    start, end := 0, len(s)-1
    for start <= end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}
