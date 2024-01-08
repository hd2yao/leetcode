package _541_Reverse_String_II

func reverseStr(s string, k int) string {
    //var result string
    //var resultSlice []string
    //
    //for i := 0; i < len(s); i += 2 * k {
    //    end := i + 2*k
    //    if end > len(s) {
    //        end = len(s)
    //    }
    //    resultSlice = append(resultSlice, s[i:end])
    //}
    //for i := 0; i < len(resultSlice)-1; i++ {
    //    result += reverseString([]byte(resultSlice[i][:k])) + resultSlice[i][k:2*k]
    //}
    //
    //if len(resultSlice[len(resultSlice)-1]) < k {
    //    reverseString([]byte(resultSlice[len(resultSlice)-1]))
    //    result += reverseString([]byte(resultSlice[len(resultSlice)-1]))
    //} else {
    //    result += reverseString([]byte(resultSlice[len(resultSlice)-1][:k])) + resultSlice[len(resultSlice)-1][k:]
    //}
    //
    //return result
    ss := []byte(s)
    for i := 0; i < len(s); i += 2 * k {
        if i+k <= len(s) {
            reverseString(ss[i : i+k])
        } else {
            reverseString(ss[i:])
        }
    }
    return string(ss)
}

func reverseString(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
