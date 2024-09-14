package day3

func isHappy(n int) bool {
    numsMap := make(map[int]struct{})
    for n != 1 {
        sum := 0
        // 取每位数字
        for n != 0 {
            nums1 := n % 10
            sum += nums1 * nums1
            n = n / 10
        }
        // 判断是否进入无限循环
        if _, ok := numsMap[sum]; ok {
            return false
        }
        numsMap[sum] = struct{}{}
        n = sum
    }
    return true
}
