package _459_Repeated_Substring_Pattern

func repeatedSubstringPattern(s string) bool {
    ss := (s + s)[1 : 2*len(s)-1]
    return strStr(ss, s)
}

func strStr(haystack string, needle string) bool {

    for i := 0; i <= len(haystack)-len(needle); i++ {
        j := 0
        if haystack[i] == needle[j] {
            index := i
            for index < len(haystack) && j < len(needle) && haystack[index] == needle[j] {
                index++
                j++
            }
            if j != len(needle) {
                continue
            }
            return true
        }
    }
    return false
}
