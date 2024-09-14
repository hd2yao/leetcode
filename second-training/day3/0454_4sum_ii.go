package day3

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    sumMap := make(map[int]int)
    count := 0
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            sumMap[-nums1[i]-nums2[j]]++
        }
    }
    for i := 0; i < len(nums3); i++ {
        for j := 0; j < len(nums4); j++ {
            if v, ok := sumMap[nums3[i]+nums4[j]]; ok {
                count += v
            }
        }
    }
    return count
}
