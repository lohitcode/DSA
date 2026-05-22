package main

import "testing"

func TestSwimInWater(t *testing.T) {
    tests := []struct {
        grid [][]int
        want int
    }{
        {[][]int{{0, 2}, {1, 3}}, 3},
        {[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
    }
    for _, tt := range tests {
        if got := swimInWater(tt.grid); got != tt.want {
            t.Errorf("swimInWater() = %v, want %v", got, tt.want)
        }
    }
}
