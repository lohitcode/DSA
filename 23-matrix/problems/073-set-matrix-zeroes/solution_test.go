package main

import "testing"

func TestSetZeroes(t *testing.T) {
    matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
    setZeroes(matrix)
    if matrix[0][1] != 0 {
        t.Errorf("setZeroes() failed")
    }
}
