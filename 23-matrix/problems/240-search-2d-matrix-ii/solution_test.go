package main

import "testing"

func TestSearchMatrix(t *testing.T) {
    matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
    if !searchMatrix(matrix, 5) {
        t.Errorf("searchMatrix() for 5 should be true")
    }
    if searchMatrix(matrix, 20) {
        t.Errorf("searchMatrix() for 20 should be false")
    }
}
