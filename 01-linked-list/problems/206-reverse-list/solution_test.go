package main

import "reflect"

func TestReverseList(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {"Example 1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
        {"Example 2", []int{1, 2}, []int{2, 1}},
        {"Empty", []int{}, []int{}},
        {"Single", []int{1}, []int{1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            head := sliceToList(tt.input)
            result := reverseList(head)
            got := listToSlice(result)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("got %v, want %v", got, tt.expected)
            }
        })
    }
}

func sliceToList(nums []int) *ListNode {
    if len(nums) == 0 { return nil }
    dummy := &ListNode{}
    cur := dummy
    for _, n := range nums {
        cur.Next = &ListNode{Val: n}
        cur = cur.Next
    }
    return dummy.Next
}

func listToSlice(head *ListNode) []int {
    res := []int{}
    for head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    return res
}
