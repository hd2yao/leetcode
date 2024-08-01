package day45

import "sort"

func findContentChildren(g []int, s []int) int {
	count := 0
	sort.Ints(g)
	sort.Ints(s)

	for i, j := 0, 0; i < len(g) && j < len(s); i, j = i+1, j+1 {
		for i < len(g) && j < len(s) {
			if g[i] > s[j] {
				j++
			} else {
				i++
				j++
				count++
			}
		}
	}

	return count
}

// 单层循环

// 用小饼干优先喂胃口小的孩子
func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	// 遍历饼干，下标控制孩子
	child := 0

	for i := 0; i < len(s) && child < len(g); i++ {
		if g[child] <= s[i] {
			child++
		}
	}

	return child
}

// 用大饼干优先喂胃口大的孩子
func findContentChildren3(g []int, s []int) int {
	count := 0
	sort.Ints(g)
	sort.Ints(s)

	// 遍历孩子，下标控制饼干
	index := len(s) - 1

	for i := len(g) - 1; i >= 0 && index >= 0; i-- {
		if g[i] <= s[index] {
			index--
			count++
		}
	}

	return count
}
