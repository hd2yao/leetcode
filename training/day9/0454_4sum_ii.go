package day9

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    result := 0
    targetMap := make(map[int]int)
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            targetMap[nums1[i]+nums2[j]]++
        }
    }

    for i := 0; i < len(nums3); i++ {
        for j := 0; j < len(nums4); j++ {
            if counts, ok := targetMap[-nums3[i]-nums4[j]]; ok {
                result += counts
            }
        }
    }

    return result
}
