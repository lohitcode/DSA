package main

import "testing"

func TestRotate(t *testing.T) {
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    rotate(matrix)
    expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] != expected[i][j] {
                t.Errorf("rotate() = %v, want %v", matrix, expected)
            }
        }
    }
}
