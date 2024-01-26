package _020_Valid_Parentheses

import (
    "fmt"
    "testing"
)

type question0020 struct {
    param0020
    ans0020
}

type param0020 struct {
    s string
}

type ans0020 struct {
    result bool
}

func TestIsValid(t *testing.T) {
    qs := []question0020{
        {
            param0020{"()"},
            ans0020{true},
        },
        {
            param0020{"()[]{}"},
            ans0020{true},
        },
        {
            param0020{"(]"},
            ans0020{false},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 20-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0020, q.param0020
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, isValid(p.s))
        fmt.Println()
    }
}
