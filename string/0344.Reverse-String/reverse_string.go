package _344_Reverse_String

func reverseString(s []byte) {
    if len(s) > 1 {
        left, right := 0, len(s)-1
        for left < right {
            s[left], s[right] = s[right], s[left]
            left++
            right--
        }
    }
}
