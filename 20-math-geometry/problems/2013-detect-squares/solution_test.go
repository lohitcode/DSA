package main

import "testing"

func TestDetectSquares(t *testing.T) {
    ds := Constructor()
    ds.Add([]int{3, 10})
    ds.Add([]int{11, 2})
    ds.Add([]int{3, 2})
    if got := ds.Count([]int{11, 10}); got != 1 {
        t.Errorf("Count() = %v, want 1", got)
    }
}
