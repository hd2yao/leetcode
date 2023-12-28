package _202_Happy_Number

import (
	"fmt"
	"testing"
)

type question0202 struct {
	param0202
	ans0202
}

type param0202 struct {
	n int
}

type ans0202 struct {
	result bool
}

func TestIsHappy(t *testing.T) {

	qs := []question0202{

		{
			param0202{19},
			ans0202{true},
		},

		{
			param0202{2},
			ans0202{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 202------------------------\n")

	for _, q := range qs {
		expected, p := q.ans0202, q.param0202
		fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, isHappy(p.n))
	}
	fmt.Printf("\n\n\n")
}
