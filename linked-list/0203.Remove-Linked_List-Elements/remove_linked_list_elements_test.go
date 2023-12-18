package _203_Remove_Linked_List_Elements

import (
	"fmt"
	"github.com/hd2yao/leetcode/structures"
	"testing"
)

type question0203 struct {
	para0203
	ans0203
}

type para0203 struct {
	nums   []int
	target int
}

type ans0203 struct {
	one []int
}

func TestRemoveElements(t *testing.T) {
	qs := []question0203{
		{
			para0203{[]int{1, 2, 6, 3, 4, 5, 6}, 6},
			ans0203{[]int{1, 2, 3, 4, 5}},
		},
		{
			para0203{[]int{}, 1},
			ans0203{[]int{}},
		},
		{
			para0203{[]int{7, 7, 7, 7}, 7},
			ans0203{[]int{}},
		},
		//{
		//	para0203{[]int{1, 2, 3, 4, 5}, 11},
		//	ans0203{[]int{}},
		//},
	}

	fmt.Printf("------------------------Leetcode Problem 203------------------------\n")

	for _, q := range qs {
		expected, p := q.ans0203, q.para0203
		fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(removeElements(structures.Ints2List(p.nums), p.target)))
		fmt.Println()
	}

}
