package day46

func maxProfit(prices []int) int {
    max := 0
    profit := make([]int, 0)

    // 构造利润序列
    for i := 1; i < len(prices); i++ {
        profit = append(profit, prices[i]-prices[i-1])
    }

    // 贪心：选择正利润
    for i := 0; i < len(profit); i++ {
        if profit[i] > 0 {
            max += profit[i]
        }
    }
    return max
}

// 简化
func maxProfit2(prices []int) int {
    max := 0
    for i := 1; i < len(prices); i++ {
        if prices[i]-prices[i-1] > 0 {
            max += prices[i] - prices[i-1]
        }
    }
    return max
}
