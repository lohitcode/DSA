package main

import "testing"

func TestUniquePaths(t *testing.T) {
    tests := []struct{ m, n, want int }{
        {3, 7, 28}, {3, 2, 3},
    }
    for _, tt := range tests {
        if got := uniquePaths(tt.m, tt.n); got != tt.want {
            t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
        }
    }
}
