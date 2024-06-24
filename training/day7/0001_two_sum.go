package day7

func twoSum(nums []int, target int) []int {
    index := make(map[int]int, len(nums))
    result := make([]int, 0)
    for i, num := range nums {
        if index, ok := index[num]; ok {
            result = append(result, index)
            result = append(result, i)
            return result
        }
        index[target-num] = i
    }
    return nil
}
