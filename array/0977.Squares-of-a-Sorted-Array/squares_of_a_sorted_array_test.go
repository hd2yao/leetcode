package _977_Squares_of_a_Sorted_Array

import (
    "fmt"
    "testing"
)

type question0977 struct {
    para0977
    ans0977
}

// para 是参数
type para0977 struct {
    nums []int
}

// ans 是答案
type ans0977 struct {
    result []int
}

func TestSortedSquares(t *testing.T) {
    qs := []question0977{
        {
            para0977{[]int{-4, -1, 0, 3, 10}},
            ans0977{[]int{0, 1, 9, 16, 100}},
        },
        {
            para0977{[]int{1}},
            ans0977{[]int{1}},
        },

        {
            para0977{[]int{-7, -3, 2, 3, 11}},
            ans0977{[]int{4, 9, 9, 49, 121}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 977------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0977, q.para0977
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, sortedSquares(p.nums))
        fmt.Println()
    }
    fmt.Printf("\n\n\n")
}
