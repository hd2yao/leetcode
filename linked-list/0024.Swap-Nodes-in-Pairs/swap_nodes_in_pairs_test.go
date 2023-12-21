package _024_Swap_Nodes_in_Pairs

import (
	"fmt"
	"github.com/hd2yao/leetcode/structures"
	"testing"
)

type question0024 struct {
	para0024
	ans0024
}

type para0024 struct {
	nums []int
}

type ans0024 struct {
	result []int
}

func TestSwapPairs(t *testing.T) {
	qs := []question0024{
		{
			para0024{[]int{1, 2, 3, 4}},
			ans0024{[]int{2, 1, 4, 3}},
		},
		{
			para0024{[]int{}},
			ans0024{[]int{}},
		},
		{
			para0024{[]int{1}},
			ans0024{[]int{1}},
		},
		{
			para0024{[]int{1, 2, 3, 4, 5}},
			ans0024{[]int{2, 1, 4, 3, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 24-------------------------\n")

	for _, q := range qs {
		expected, p := q.ans0024, q.para0024
		fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(swapPairs(structures.Ints2List(p.nums))))
		fmt.Println()
	}
}
