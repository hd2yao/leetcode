package _541_Reverse_String_II

import (
	"fmt"
	"testing"
)

type question0541 struct {
	param0541
	ans0541
}

type param0541 struct {
	s string
	k int
}

type ans0541 struct {
	result string
}

func TestReverseStr(t *testing.T) {

	qs := []question0541{
		{
			param0541{"abcdefg", 2},
			ans0541{"bacdfeg"},
		},

		{
			param0541{"abcd", 2},
			ans0541{"bacd"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 541------------------------\n")

	for _, q := range qs {
		expected, p := q.ans0541, q.param0541
		fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p.s, expected, reverseStr(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
