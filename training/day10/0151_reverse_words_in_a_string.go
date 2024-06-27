package day10

func reverseWords(s string) string {
    b := []byte(s)
    slow := 0
    // 快慢指针法，去除头部以及单词间多余的空格
    for i := 0; i < len(b); i++ {
        if b[i] != ' ' {
            if slow != 0 {
                b[slow] = ' '
                slow++
            }
            for i < len(b) && b[i] != ' ' {
                b[slow] = b[i]
                i++
                slow++
            }
        }
    }
    // 去除尾部的多余空格
    b = b[:slow]

    // 先整体反转
    reverse1(b)

    // 每个单词反转
    // start 用于记录每个单词的起始位置
    start := 0
    for i := 0; i < len(b); i++ {
        if b[i] == ' ' {
            reverse1(b[start:i])
            start = i + 1
        }
    }
    // 单独翻转最后一个单词
    reverse1(b[start:])

    return string(b)
}

func reverse1(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
