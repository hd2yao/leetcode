package _028_Find_the_Index_of_the_First_Occurrence_in_a_String

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	index := 0
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		if haystack[i] == needle[j] {
			index = i
			for index < len(haystack) && j < len(needle) && haystack[index] == needle[j] {
				index++
				j++
			}
			if j != len(needle) {
				continue
			}
			return i
		}
	}
	return -1
}
