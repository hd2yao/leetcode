package _707_Design_Linked_List

import (
    "fmt"
    "testing"
)

func TestDesignSingleLinkedList(t *testing.T) {
    obj := Constructor()
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    param1 := obj.Get(1)
    fmt.Printf("param_1 = %v obj = %v\n", param1, List2Ints(&obj))
    obj.AddAtHead(38)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtHead(45)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.DeleteAtIndex(2)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtIndex(1, 24)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtTail(36)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtIndex(3, 72)
    obj.AddAtTail(76)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtHead(7)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtHead(36)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtHead(34)
    fmt.Printf("obj = %v\n", List2Ints(&obj))
    obj.AddAtTail(91)
    fmt.Printf("obj = %v\n", List2Ints(&obj))

    fmt.Printf("\n\n\n")

    obj1 := Constructor()
    fmt.Printf("obj1 = %v\n", List2Ints(&obj1))
    param2 := obj1.Get(0)
    fmt.Printf("param_2 = %v obj1 = %v\n", param2, List2Ints(&obj1))
    obj1.AddAtIndex(1, 2)
    fmt.Printf("obj1 = %v\n", List2Ints(&obj1))
    param2 = obj1.Get(0)
    fmt.Printf("param_2 = %v obj1 = %v\n", param2, List2Ints(&obj1))
    param2 = obj1.Get(1)
    fmt.Printf("param_2 = %v obj1 = %v\n", param2, List2Ints(&obj1))
    obj1.AddAtIndex(0, 1)
    fmt.Printf("obj1 = %v\n", List2Ints(&obj1))
    param2 = obj1.Get(0)
    fmt.Printf("param_1 = %v obj1 = %v\n", param2, List2Ints(&obj1))
    param2 = obj1.Get(1)
    fmt.Printf("param_2 = %v obj1 = %v\n", param2, List2Ints(&obj1))
}

func List2Ints(head *MyLinkedList) []int {
    res := []int{}
    cur := head.DummyHead.Next
    for cur != nil {
        res = append(res, cur.Val)
        cur = cur.Next
    }
    return res
}
