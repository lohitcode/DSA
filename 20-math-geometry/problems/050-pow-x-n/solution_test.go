package main

import "testing"

func TestMyPow(t *testing.T) {
    tests := []struct {
        x, n float64
        want float64
    }{
        {2.0, 10, 1024.0},
        {2.1, 3, 9.261},
        {2.0, -2, 0.25},
    }
    for _, tt := range tests {
        got := myPow(tt.x, int(tt.n))
        if absFloat(got - tt.want) > 0.00001 {
            t.Errorf("myPow() = %v, want %v", got, tt.want)
        }
    }
}

func absFloat(x float64) float64 { if x < 0 { return -x }; return x }
