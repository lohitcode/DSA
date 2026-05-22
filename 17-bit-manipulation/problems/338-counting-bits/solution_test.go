package main

import "testing"

func TestCountBits(t *testing.T) {
    tests := []struct {
        n    int
        want []int
    }{
        {2, []int{0, 1, 1}},
        {5, []int{0, 1, 1, 2, 1, 2}},
    }
    for _, tt := range tests {
        if got := countBits(tt.n); len(got) != len(tt.want) {
            t.Errorf("countBits() length mismatch")
        } else {
            for i := range got {
                if got[i] != tt.want[i] {
                    t.Errorf("countBits()[%d] = %v, want %v", i, got[i], tt.want[i])
                }
            }
        }
    }
}
