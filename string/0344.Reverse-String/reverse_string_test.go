package _344_Reverse_String

import (
    "fmt"
    "testing"
)

type question0344 struct {
    param0344
    ans0344
}

type param0344 struct {
    strings []byte
}

type ans0344 struct {
    result []byte
}

func TestTwoSum(t *testing.T) {

    qs := []question0344{
        {
            param0344{[]byte{'h', 'e', 'l', 'l', 'o'}},
            ans0344{[]byte{'o', 'l', 'l', 'e', 'h'}},
        },

        {
            param0344{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
            ans0344{[]byte{'h', 'a', 'n', 'n', 'a', 'H'}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 344------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0344, q.param0344
        reverseString(p.strings)
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p.strings, expected, p.strings)
    }
    fmt.Printf("\n\n\n")
}
