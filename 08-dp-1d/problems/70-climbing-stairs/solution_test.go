package main

import "testing"

func TestClimbStairs(t *testing.T) {
    tests := []struct{ n, want int }{
        {2, 2}, {3, 3}, {4, 5}, {5, 8},
    }
    for _, tt := range tests {
        if got := climbStairs(tt.n); got != tt.want {
            t.Errorf("climbStairs() = %v, want %v", got, tt.want)
        }
    }
}
