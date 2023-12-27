package _349_Intersection_of_Two_Arrays

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, len(nums1)+len(nums2))
	for _, value := range nums1 {
		if _, ok := numMap[value]; !ok {
			numMap[value]++
		}
	}
	for _, value := range nums2 {
		if _, ok := numMap[value]; ok {
			numMap[value]--
		}
	}
	result := []int{}
	for value, count := range numMap {
		if count <= 0 {
			result = append(result, value)
		}
	}
	return result
}
