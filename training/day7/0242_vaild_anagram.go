package day7

// map
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapStr := make(map[string]int, len(s))
    for _, str := range s {
        mapStr[string(str)]++
    }
    for _, str := range t {
        num, ok := mapStr[string(str)]
        if ok && num > 0 {
            mapStr[string(str)]--
        } else {
            return false
        }
    }
    return true
}

// array
func isAnagramArray(s string, t string) bool {
    array := [26]int{}

    for _, v := range s {
        array[v-97]++
    }

    for _, v := range t {
        array[v-97]--
    }

    for _, num := range array {
        if num != 0 {
            return false
        }
    }
    return true
}
