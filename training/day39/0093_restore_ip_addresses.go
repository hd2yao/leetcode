package day39

import (
    "strconv"
    "strings"
)

func restoreIpAddresses(s string) []string {
    if len(s) < 4 || len(s) > 12 {
        return nil
    }

    result := make([]string, 0)
    path := ""

    var backtracking func(s, path string, dot int)
    backtracking = func(s, path string, dot int) {
        if dot == 3 {
            if isValid(s) {
                path += s
                result = append(result, path)
                path = ""
            }
            return
        }

        for i := 0; i < len(s); i++ {
            ip := s[:i+1]
            if !isValid(ip) {
                continue
            }
            path += ip + "."
            backtracking(s[i+1:], path, dot+1)
            path = path[:len(path)-1]
            path = path[:strings.LastIndex(path, ".")+1]
        }
    }

    backtracking(s, path, 0)
    return result
}

func isValid(s string) bool {
    if (len(s) > 1 && s[0] == '0') || len(s) > 3 {
        return false
    }
    num, err := strconv.Atoi(s)
    if err != nil || num > 255 {
        return false
    }
    return true
}
