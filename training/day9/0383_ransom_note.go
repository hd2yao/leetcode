package day9

func canConstruct(ransomNote string, magazine string) bool {
    alphaMap := [26]int{}
    for _, a := range magazine {
        alphaMap[a-'a']++
    }
    for _, b := range ransomNote {
        alphaMap[b-'a']--
        if alphaMap[b-'a'] < 0 {
            return false
        }
    }
    return true
}
