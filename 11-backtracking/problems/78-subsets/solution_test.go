package main

import "testing"

func TestSubsets(t *testing.T) {
    got := subsets([]int{1, 2, 3})
    if len(got) != 8 {
        t.Errorf("subsets() length = %v, want 8", len(got))
    }
}
