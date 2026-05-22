package main

import "testing"

func TestPlusOne(t *testing.T) {
    tests := []struct {
        digits, want []int
    }{
        {[]int{1, 2, 3}, []int{1, 2, 4}},
        {[]int{9}, []int{1, 0}},
        {[]int{9, 9}, []int{1, 0, 0}},
    }
    for _, tt := range tests {
        got := plusOne(tt.digits)
        if len(got) != len(tt.want) {
            t.Errorf("plusOne() length mismatch")
        } else {
            for i := range got {
                if got[i] != tt.want[i] {
                    t.Errorf("plusOne()[%d] = %v, want %v", i, got[i], tt.want[i])
                }
            }
        }
    }
}
