package main

import "testing"

func TestSearch(t *testing.T) {
    tests := []struct{ nums []int; target, want int }{
        {[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
        {[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
    }
    for _, tt := range tests {
        if got := search(tt.nums, tt.target); got != tt.want {
            t.Errorf("search() = %v, want %v", got, tt.want)
        }
    }
}
