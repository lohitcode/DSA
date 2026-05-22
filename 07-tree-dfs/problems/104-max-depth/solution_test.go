package main

import "testing"

func TestMaxDepth(t *testing.T) {
    tests := []struct{ tree *TreeNode; want int }{
        {&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 3},
        {nil, 0},
        {&TreeNode{Val: 1}, 1},
    }
    for _, tt := range tests {
        if got := maxDepth(tt.tree); got != tt.want {
            t.Errorf("maxDepth() = %v, want %v", got, tt.want)
        }
    }
}
