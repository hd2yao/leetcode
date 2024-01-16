package _151_Reverse_Words_in_a_String

func reverseWords(s string) string {
    ss := []byte(s)
    j := 0
    for i := 0; i < len(ss); i++ {
        if ss[i] != ' ' {
            if j != 0 {
                ss[j] = ' '
                j++
            }
            for i < len(ss) && ss[i] != ' ' {
                ss[j] = ss[i]
                j++
                i++
            }
        }
    }
    ss = ss[:j]

    reverseStr(ss)

    index := 0
    for i := 0; i <= len(ss); i++ {
        if i == len(ss) || ss[i] == ' ' {
            reverseStr(ss[index:i])
            index = i + 1
        }
    }
    return string(ss)
}

func reverseStr(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
