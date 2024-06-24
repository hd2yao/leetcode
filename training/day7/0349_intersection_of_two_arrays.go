package day7

func intersection(nums1 []int, nums2 []int) []int {
    numMap := make(map[int]int, len(nums1))
    result := make([]int, 0)
    for _, num := range nums1 {
        if _, ok := numMap[num]; !ok {
            numMap[num] = 1
        }
    }
    for _, num := range nums2 {
        if _, ok := numMap[num]; ok {
            result = append(result, num)
            delete(numMap, num)
        }
    }
    return result
}
