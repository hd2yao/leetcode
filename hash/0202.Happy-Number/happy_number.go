package _202_Happy_Number

func isHappy(n int) bool {
    numMap := make(map[int]int)
    nums := sumFunc(n, 0)
    for nums != 1 {
        if _, ok := numMap[nums]; !ok {
            numMap[nums] = 1
            nums = sumFunc(nums, 0)
        } else {
            return false
        }
    }
    return true
}

func sumFunc(num, sum int) int {
    if num/10 == 0 {
        return num*num + sum
    }
    tmp := num % 10
    sum += tmp * tmp
    res := num / 10
    return sumFunc(res, sum)
}
