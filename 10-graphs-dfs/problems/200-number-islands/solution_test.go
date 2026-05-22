package main

import "testing"

func TestNumIslands(t *testing.T) {
    grid := [][]byte{
        {'1','1','1','1','0'},
        {'1','1','0','1','0'},
        {'1','1','0','0','0'},
        {'0','0','0','0','0'},
    }
    want := 1
    if got := numIslands(grid); got != want {
        t.Errorf("numIslands() = %v, want %v", got, want)
    }
}
