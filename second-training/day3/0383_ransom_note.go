package day3

func canConstruct(ransomNote string, magazine string) bool {
    alphaArr := [26]int{}
    for i := 0; i < len(magazine); i++ {
        alphaArr[magazine[i]-'a']++
    }
    for i := 0; i < len(ransomNote); i++ {
        if alphaArr[ransomNote[i]-'a'] == 0 {
            return false
        }
        alphaArr[ransomNote[i]-'a']--
    }
    return true
}
