package main

import "testing"

func TestFindOrder(t *testing.T) {
    tests := []struct {
        numCourses    int
        prerequisites [][]int
        wantLen       int
    }{
        {4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, 4},
        {1, [][]int{}, 1},
    }
    for _, tt := range tests {
        got := findOrder(tt.numCourses, tt.prerequisites)
        if len(got) != tt.wantLen {
            t.Errorf("findOrder() length = %v, want %v", len(got), tt.wantLen)
        }
    }
}
