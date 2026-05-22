package main

import "testing"

func TestCanFinish(t *testing.T) {
    tests := []struct {
        numCourses    int
        prerequisites [][]int
        want          bool
    }{
        {2, [][]int{{1, 0}}, true},
        {2, [][]int{{1, 0}, {0, 1}}, false},
        {5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
    }
    for _, tt := range tests {
        if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
            t.Errorf("canFinish() = %v, want %v", got, tt.want)
        }
    }
}
