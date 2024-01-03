package _383_Ransom_Note

import (
    "fmt"
    "testing"
)

type question0383 struct {
    param0383
    ans0383
}

type param0383 struct {
    ransomNote string
    magazine   string
}

type ans0383 struct {
    result bool
}

func TestCanConstruct(t *testing.T) {

    qs := []question0383{

        {
            param0383{"a", "b"},
            ans0383{false},
        },

        {
            param0383{"aa", "ab"},
            ans0383{false},
        },

        {
            param0383{"aa", "aab"},
            ans0383{true},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 383------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0383, q.param0383
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, canConstruct(p.ransomNote, p.magazine))
    }
    fmt.Printf("\n\n\n")
}
