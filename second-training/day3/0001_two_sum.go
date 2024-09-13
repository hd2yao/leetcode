package day3

func twoSum(nums []int, target int) []int {
    sumMap := make(map[int]int)
    result := make([]int, 0)
    for i, num := range nums {
        if index, ok := sumMap[num]; ok {
            result = append(result, i, index)
            break
        }
        sumMap[target-num] = i
    }
    return result
}
