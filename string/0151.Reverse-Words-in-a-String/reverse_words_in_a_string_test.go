package _151_Reverse_Words_in_a_String

import (
    "fmt"
    "testing"
)

type question0151 struct {
    param0151
    ans0151
}

type param0151 struct {
    s string
}

type ans0151 struct {
    result string
}

func TestReverseWords(t *testing.T) {

    qs := []question0151{
        {
            param0151{"the sky is blue"},
            ans0151{"blue is sky the"},
        },

        {
            param0151{"  hello world  "},
            ans0151{"world hello"},
        },

        {
            param0151{"a good   example"},
            ans0151{"example good a"},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 151------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0151, q.param0151
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p.s, expected, reverseWords(p.s))
    }
    fmt.Printf("\n\n\n")
}
