package _383_Ransom_Note

// 使用 map
func canConstruct(ransomNote string, magazine string) bool {
    magazineMap := make(map[byte]int, len(magazine))
    for _, a := range magazine {
        magazineMap[byte(a)]++
    }
    for _, b := range ransomNote {
        counts, ok := magazineMap[byte(b)]
        if counts == 0 || !ok {
            return false
        }
        magazineMap[byte(b)]--
    }
    return true
}

// 因为都是小写字母，因此可以使用 [26]int 数组
func canConstructArray(ransomNote string, magazine string) bool {
    magazineArray := [26]int{}
    for _, a := range magazine {
        magazineArray[a-'a']++
    }
    for _, b := range ransomNote {
        counts := magazineArray[b-'a']
        if counts <= 0 {
            return false
        }
        magazineArray[b-'a']--
    }
    return true
}
