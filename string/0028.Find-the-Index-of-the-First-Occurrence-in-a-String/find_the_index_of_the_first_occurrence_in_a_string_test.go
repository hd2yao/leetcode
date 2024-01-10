package _028_Find_the_Index_of_the_First_Occurrence_in_a_String

import (
    "fmt"
    "testing"
)

type question0028 struct {
    param0028
    ans0028
}

type param0028 struct {
    haystack string
    needle   string
}

type ans0028 struct {
    result int
}

func TestStrStr(t *testing.T) {

    qs := []question0028{
        {
            param0028{"sadbutsad", "sad"},
            ans0028{0},
        },

        {
            param0028{"leetcode", "leeto"},
            ans0028{-1},
        },

        {
            param0028{"aaa", "aaa"},
            ans0028{0},
        },

        {
            param0028{"mississippi", "issip"},
            ans0028{4},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 28-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0028, q.param0028
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, strStr(p.haystack, p.needle))
    }
    fmt.Printf("\n\n\n")
}
