package day4

func reverseStr(s string, k int) string {
    ss := []byte(s)
    for i := 0; i < len(ss); i = i + 2*k {
        if i+k < len(ss) {
            reverse(ss[i : i+k])
        } else {
            reverse(ss[i:])
        }
    }
    return string(ss)
}

func reverse(s []byte) {
    for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}
