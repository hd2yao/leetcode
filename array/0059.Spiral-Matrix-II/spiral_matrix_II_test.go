package _059_Spiral_Matrix_II

import (
    "fmt"
    "testing"
)

type question0059 struct {
    para0059
    ans0059
}

type para0059 struct {
    num int
}

type ans0059 struct {
    one [][]int
}

func TestGenerateMatrix(t *testing.T) {
    qs := []question0059{
        {
            para0059{3},
            ans0059{[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
        },
        {
            para0059{1},
            ans0059{[][]int{{1}}},
        },
        {
            para0059{4},
            ans0059{[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 59-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0059, q.para0059
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, generateMatrix(p.num))
        fmt.Println()
    }

}
