package _242_Vaild_Anagram

import (
    "fmt"
    "testing"
)

type question0242 struct {
    param0242
    ans0242
}

type param0242 struct {
    s string
    t string
}

type ans0242 struct {
    result bool
}

func TestIsAnagram(t *testing.T) {

    qs := []question0242{

        {
            param0242{"", ""},
            ans0242{true},
        },

        {
            param0242{"", "1"},
            ans0242{false},
        },

        {
            param0242{"anagram", "nagaram"},
            ans0242{true},
        },

        {
            param0242{"rat", "car"},
            ans0242{false},
        },

        {
            param0242{"aa", "bb"},
            ans0242{false},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 242------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0242, q.param0242
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, isAnagramArray(p.s, p.t))
    }
    fmt.Printf("\n\n\n")
}
