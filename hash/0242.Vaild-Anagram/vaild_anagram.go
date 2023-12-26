package _242_Vaild_Anagram

// 使用 map
func isAnagramMap(s string, t string) bool {
	alphHash := make(map[rune]int, 26)
	for _, a := range s {
		alphHash[a]++
	}

	for _, b := range t {
		if _, ok := alphHash[b]; !ok {
			return false
		}
		alphHash[b]--
	}
	for _, count := range alphHash {
		if count != 0 {
			return false
		}
	}
	return true
}

// 使用数组
func isAnagramArray(s string, t string) bool {
	alphaArray := [26]int{}
	for _, a := range s {
		alphaArray[a-'a']++
	}
	for _, b := range t {
		alphaArray[b-'a']--
	}
	for _, value := range alphaArray {
		if value != 0 {
			return false
		}
	}
	return true
}
