package day3

func intersection(nums1 []int, nums2 []int) []int {
    numsMap := make(map[int]bool)
    result := make([]int, 0)
    for _, i := range nums1 {
        numsMap[i] = true
    }
    for _, i := range nums2 {
        if v, ok := numsMap[i]; ok && v {
            result = append(result, i)
            numsMap[i] = false
        }
    }
    return result
}
