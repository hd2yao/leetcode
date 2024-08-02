package day46

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	minStep := 0
	curDistance := 0
	nextDistance := 0

	for i := 0; i <= curDistance; i++ {

		if i+nums[i] > nextDistance {
			nextDistance = i + nums[i]
		}

		if i == curDistance {
			minStep++
			curDistance = nextDistance
			if nextDistance >= len(nums)-1 {
				break
			}
		}
	}

	return minStep
}
