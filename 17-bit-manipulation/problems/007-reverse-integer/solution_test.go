package main

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        x, want int
    }{
        {123, 321},
        {-123, -321},
        {120, 21},
        {0, 0},
        {1534236469, 0}, // overflow
    }
    for _, tt := range tests {
        if got := reverse(tt.x); got != tt.want {
            t.Errorf("reverse() = %v, want %v", got, tt.want)
        }
    }
}
