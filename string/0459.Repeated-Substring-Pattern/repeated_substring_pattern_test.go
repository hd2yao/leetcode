package _459_Repeated_Substring_Pattern

import (
    "fmt"
    "testing"
)

type question0459 struct {
    param0459
    ans0459
}

type param0459 struct {
    s string
}

type ans0459 struct {
    result bool
}

func TestStrStr(t *testing.T) {

    qs := []question0459{
        {
            param0459{"abab"},
            ans0459{true},
        },

        {
            param0459{"aba"},
            ans0459{false},
        },

        {
            param0459{"abcabcabcabc"},
            ans0459{true},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 459------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0459, q.param0459
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, repeatedSubstringPattern(p.s))
    }
    fmt.Printf("\n\n\n")
}
