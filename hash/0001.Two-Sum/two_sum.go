package _001_Two_Sum

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := targetMap[nums[i]]; !ok {
			targetMap[target-nums[i]] = i
		} else {
			return []int{i, targetMap[nums[i]]}
		}
	}
	return nil
}
