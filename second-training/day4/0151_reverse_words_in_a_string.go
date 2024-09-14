package day4

func reverseWords(s string) string {
    ss := []byte(s)
    // 1. 先去除多余的空格
    slow := 0
    for i := 0; i < len(ss); i++ {
        if ss[i] != ' ' {
            if slow != 0 {
                ss[slow] = ' '
                slow++
            }
            for i < len(ss) && ss[i] != ' ' {
                ss[slow] = ss[i]
                i++
                slow++
            }
        }
    }
    ss = ss[:slow]

    // 2. 整体翻转
    reverse(ss)

    // 3. 每个单词再翻转
    start := 0
    for i := 0; i < len(ss); i++ {
        if ss[i] == ' ' {
            reverse(ss[start:i])
            start = i + 1
        }
    }
    reverse(ss[start:])

    return string(ss)
}
