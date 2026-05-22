package main

import "testing"

func TestKthSmallest(t *testing.T) {
    tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
    if got := kthSmallest(tree, 1); got != 1 {
        t.Errorf("kthSmallest() = %v, want 1", got)
    }
}
