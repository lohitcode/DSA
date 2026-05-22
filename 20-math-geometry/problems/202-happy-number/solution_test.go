package main

import "testing"

func TestIsHappy(t *testing.T) {
    tests := []struct {
        n    int
        want bool
    }{
        {19, true},
        {2, false},
        {1, true},
    }
    for _, tt := range tests {
        if got := isHappy(tt.n); got != tt.want {
            t.Errorf("isHappy() = %v, want %v", got, tt.want)
        }
    }
}
