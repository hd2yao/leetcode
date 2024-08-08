package day52

import "strconv"

func monotoneIncreasingDigits(n int) int {
    strNum := strconv.Itoa(n)
    byteNum := []byte(strNum)

    mark := len(byteNum)
    for i := len(byteNum) - 1; i > 0; i-- {
        if byteNum[i] >= byteNum[i-1] {
            continue
        }
        byteNum[i-1]--
        mark = i
    }

    for i := mark; i < len(byteNum); i++ {
        byteNum[i] = '9'
    }

    num, _ := strconv.Atoi(string(byteNum))
    return num
}
