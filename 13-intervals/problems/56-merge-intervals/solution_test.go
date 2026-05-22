package main

import "testing"

func TestMerge(t *testing.T) {
    got := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
    if len(got) != 3 {
        t.Errorf("merge() length = %v, want 3", len(got))
    }
}
