package main

import "testing"

func TestSetZeroes(t *testing.T) {
    matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
    setZeroes(matrix)
    expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] != expected[i][j] {
                t.Errorf("setZeroes() = %v, want %v", matrix, expected)
            }
        }
    }
}
