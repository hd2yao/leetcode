package _151_Reverse_Words_in_a_String

func reverseWords(s string) string {
	b := []byte(s)
	// 1. 快慢指针：移除前面、中间、后面多余的空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[:slow]

	// 2. 反转整个字符串
	reverseString(b)

	// 3. 反转每个单词
	j := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || i == len(b) {
			reverseString(b[j:i])
			j = i + 1
		}
	}
	return string(b)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
