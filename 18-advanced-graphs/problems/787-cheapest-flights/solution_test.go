package main

import "testing"

func TestFindCheapestPrice(t *testing.T) {
    tests := []struct {
        n       int
        flights [][]int
        src     int
        dst     int
        k       int
        want    int
    }{
        {4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1, 700},
        {3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
    }
    for _, tt := range tests {
        if got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k); got != tt.want {
            t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
        }
    }
}
