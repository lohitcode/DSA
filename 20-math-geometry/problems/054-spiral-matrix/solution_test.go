package main

import "testing"

func TestSpiralOrder(t *testing.T) {
    tests := []struct {
        matrix [][]int
        want   []int
    }{
        {[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
    }
    for _, tt := range tests {
        got := spiralOrder(tt.matrix)
        if len(got) != len(tt.want) {
            t.Errorf("spiralOrder() length mismatch")
        } else {
            for i := range got {
                if got[i] != tt.want[i] {
                    t.Errorf("spiralOrder()[%d] = %v, want %v", i, got[i], tt.want[i])
                }
            }
        }
    }
}
