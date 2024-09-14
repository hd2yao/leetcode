package day3

func isAnagram(s string, t string) bool {
    result := [26]int{}
    for i := 0; i < len(s); i++ {
        result[s[i]-'a']++
    }
    for i := 0; i < len(t); i++ {
        result[t[i]-'a']--
        if result[t[i]-'a'] < 0 {
            return false
        }
    }
    if result != [26]int{} {
        return false
    }
    return true
}
