package day10

func reverseStr(s string, k int) string {
    ss := []byte(s)
    for i := 0; i < len(s); i += 2 * k {
        if i+k <= len(s) {
            reverse(ss[i : i+k])
        } else {
            reverse(ss[i:])
        }
    }
    return string(ss)
}

func reverse(s []byte) string {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
    return string(s)
}
