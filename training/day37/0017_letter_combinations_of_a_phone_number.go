package day37

func letterCombinations(digits string) []string {
    if digits == "" {
        return nil
    }
    result := make([]string, 0)
    path := ""
    letter := map[byte][]string{
        '2': {"a", "b", "c"},
        '3': {"d", "e", "f"},
        '4': {"g", "h", "i"},
        '5': {"j", "k", "l"},
        '6': {"m", "n", "o"},
        '7': {"p", "q", "r", "s"},
        '8': {"t", "u", "v"},
        '9': {"w", "x", "y", "z"},
    }

    var backtracking func(digits string, startIndex int)
    backtracking = func(digits string, startIndex int) {
        if len(path) == len(digits) {
            result = append(result, path)
            return
        }

        for i := startIndex; i < len(digits); i++ {
            for j := 0; j < len(letter[digits[startIndex]]); j++ {
                path += letter[digits[startIndex]][j]
                backtracking(digits, i+1)
                path = path[:len(path)-1]
            }
        }
    }

    backtracking(digits, 0)
    return result
}
