package _454_4Sum_II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    resultMap := make(map[int]int)
    counts := 0
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums1); j++ {
            resultMap[nums1[i]+nums2[j]]++
        }
    }

    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums1); j++ {
            if _, ok := resultMap[-(nums3[i] + nums4[j])]; ok {
                counts += resultMap[-(nums3[i] + nums4[j])]
            }
        }
    }

    return counts
}
