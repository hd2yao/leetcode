package _209_Minimum_Size_Subarray_Sum

import (
    "fmt"
    "testing"
)

type question0209 struct {
    para0209
    ans0209
}

type para0209 struct {
    nums   []int
    target int
}

type ans0209 struct {
    one int
}

func TestMinSubArrayLen(t *testing.T) {
    qs := []question0209{
        {
            para0209{[]int{2, 3, 1, 2, 4, 3}, 7},
            ans0209{2},
        },
        {
            para0209{[]int{1, 4, 4}, 4},
            ans0209{1},
        },
        {
            para0209{[]int{1, 1, 1, 1, 1, 1}, 11},
            ans0209{0},
        },
        {
            para0209{[]int{1, 2, 3, 4, 5}, 11},
            ans0209{3},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 209------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0209, q.para0209
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, minSubArrayLenOptimize(p.target, p.nums))
        fmt.Println()
    }

}
