package day49

func candy(ratings []int) int {
	count := make([]int, len(ratings))
	sum := 0

	count[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			count[i] = count[i-1] + 1
		} else {
			count[i] = 1
		}
	}

	for i := len(count) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && count[i-1] <= count[i] {
			count[i-1] = count[i] + 1
		}
	}

	for i := 0; i < len(count); i++ {
		sum += count[i]
	}

	return sum
}
