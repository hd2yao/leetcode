package _018_4Sum

import (
    "fmt"
    "testing"
)

type question0018 struct {
    param0018
    ans0018
}

type param0018 struct {
    nums   []int
    target int
}

type ans0018 struct {
    result [][]int
}

func TestTwoSum(t *testing.T) {

    qs := []question0018{
        {
            param0018{[]int{1, 0, -1, 0, -2, 2}, 0},
            ans0018{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
        },

        {
            param0018{[]int{2, 2, 2, 2, 2}, 8},
            ans0018{[][]int{{2, 2, 2, 2}}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 18-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0018, q.param0018
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, fourSum(p.nums, p.target))
    }
    fmt.Printf("\n\n\n")
}
