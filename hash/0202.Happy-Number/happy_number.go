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

// 代码精简

func isHappyOptimize(n int) bool {
    numMap := make(map[int]bool)
    for n != 1 && !numMap[n] {
        n, numMap[n] = getSum(n), true
    }
    return n == 1
}

func getSum(n int) int {
    sum := 0
    for n > 0 {
        sum += (n % 10) * (n % 10)
        n = n / 10
    }
    return sum
}
