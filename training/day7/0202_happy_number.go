package day7

func isHappy(n int) bool {
    numMap := make(map[int]struct{}, 0)
    for n != 1 {
        if _, ok := numMap[n]; !ok {
            numMap[n] = struct{}{}
            n = sum2(n)
        } else {
            return false
        }
    }
    return true
}

func sum2(n int) int {
    sum := 0
    for n/10 != 0 {
        tmp := n % 10
        n = n / 10
        sum += tmp * tmp
    }
    sum += (n % 10) * (n % 10)
    return sum
}
