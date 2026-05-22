package main

import "testing"

func TestSpiralOrder(t *testing.T) {
    got := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
    want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
    if len(got) != len(want) {
        t.Errorf("spiralOrder() length mismatch")
    }
}
