package main

import "testing"

func TestIsValidBST(t *testing.T) {
    tests := []struct {
        name string
        tree *TreeNode
        want bool
    }{
        {"valid", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, true},
        {"invalid", &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}, false},
    }
    for _, tt := range tests {
        if got := isValidBST(tt.tree); got != tt.want {
            t.Errorf("isValidBST() = %v, want %v", got, tt.want)
        }
    }
}
