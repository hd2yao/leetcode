package _015_3Sum

import (
    "fmt"
    "testing"
)

type question0015 struct {
    param0015
    ans0015
}

type param0015 struct {
    nums []int
}

type ans0015 struct {
    result [][]int
}

func TestTwoSum(t *testing.T) {

    qs := []question0015{
        {
            param0015{[]int{-1, 0, 1, 2, -1, -4}},
            ans0015{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
        },

        {
            param0015{[]int{0, 1, 1}},
            ans0015{[][]int{}},
        },

        {
            param0015{[]int{0, 0, 0}},
            ans0015{[][]int{{0, 0, 0}}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 15-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0015, q.param0015
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, threeSum(p.nums))
    }
    fmt.Printf("\n\n\n")
}
