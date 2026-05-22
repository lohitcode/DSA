package main

import "testing"

func TestGetSum(t *testing.T) {
    tests := []struct {
        a, b int
        want int
    }{
        {1, 2, 3},
        {2, 3, 5},
        {-1, 1, 0},
    }
    for _, tt := range tests {
        if got := getSum(tt.a, tt.b); got != tt.want {
            t.Errorf("getSum() = %v, want %v", got, tt.want)
        }
    }
}
