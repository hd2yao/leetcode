package _704_Binary_Search

import (
    "fmt"
    "testing"
)

type question0704 struct {
    para0704
    ans0704
}

// para 是参数
type para0704 struct {
    nums   []int
    target int
}

// ans 是答案
type ans0704 struct {
    result int
}

func TestSearch(t *testing.T) {
    qs := []question0704{
        {
            para0704{[]int{-1, 0, 3, 5, 9, 12}, 9},
            ans0704{4},
        },

        {
            para0704{[]int{-1, 0, 3, 5, 9, 12}, 2},
            ans0704{-1},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 704------------------------\n")

    for _, q := range qs {
        _, p := q.ans0704, q.para0704
        fmt.Printf("【input】:%v       【output】:%v\n", p, search(p.nums, p.target))
    }
    fmt.Printf("\n\n\n")
}

func TestSearchNoRight(t *testing.T) {
    qs := []question0704{
        {
            para0704{[]int{-1, 0, 3, 5, 9, 12}, 9},
            ans0704{4},
        },

        {
            para0704{[]int{-1, 0, 3, 5, 9, 12}, 2},
            ans0704{-1},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 704------------------------\n")

    for _, q := range qs {
        _, p := q.ans0704, q.para0704
        fmt.Printf("【input】:%v       【output】:%v\n", p, searchNoRight(p.nums, p.target))
    }
    fmt.Printf("\n\n\n")
}
