package _001_Two_Sum

import (
    "fmt"
    "testing"
)

type question0001 struct {
    param0001
    ans0001
}

type param0001 struct {
    nums   []int
    target int
}

type ans0001 struct {
    result []int
}

func TestTwoSum(t *testing.T) {

    qs := []question0001{

        {
            param0001{[]int{2, 7, 11, 15}, 9},
            ans0001{[]int{0, 1}},
        },

        {
            param0001{[]int{3, 2, 4}, 6},
            ans0001{[]int{1, 2}},
        },

        {
            param0001{[]int{3, 3}, 6},
            ans0001{[]int{0, 1}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 1--------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0001, q.param0001
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, twoSum(p.nums, p.target))
    }
    fmt.Printf("\n\n\n")
}
