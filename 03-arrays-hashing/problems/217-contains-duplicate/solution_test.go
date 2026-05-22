package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
    tests := []struct{ nums []int; want bool }{
        {[]int{1, 2, 3, 1}, true},
        {[]int{1, 2, 3, 4}, false},
        {[]int{1}, false},
    }
    for _, tt := range tests {
        if got := containsDuplicate(tt.nums); got != tt.want {
            t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
        }
    }
}
