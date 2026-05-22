package main

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
    tests := []struct {
        points [][]int
        want   int
    }{
        {[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
        {[][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}, 4},
    }
    for _, tt := range tests {
        if got := minCostConnectPoints(tt.points); got != tt.want {
            t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
        }
    }
}
