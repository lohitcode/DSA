package main

import "testing"

func TestBuildTree(t *testing.T) {
    root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
    if root.Val != 3 {
        t.Errorf("buildTree() root = %v, want 3", root.Val)
    }
}
