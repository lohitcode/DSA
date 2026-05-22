package main

import "testing"

func TestLevelOrder(t *testing.T) {
    tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
    want := [][]int{{3}, {9, 20}, {15, 7}}
    if got := levelOrder(tree); len(got) != len(want) {
        t.Errorf("levelOrder() length mismatch")
    }
}
