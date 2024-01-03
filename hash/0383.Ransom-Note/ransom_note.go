package _383_Ransom_Note

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
