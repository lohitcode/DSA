package main

import "testing"

func TestKthLargest(t *testing.T) {
    kl := Constructor(3, []int{4, 5, 8, 2})
    if got := kl.Add(3); got != 4 {
        t.Errorf("Add() = %v, want 4", got)
    }
}
