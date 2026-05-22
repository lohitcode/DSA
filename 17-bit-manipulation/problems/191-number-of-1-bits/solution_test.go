package main

import "testing"

func TestHammingWeight(t *testing.T) {
    tests := []struct {
        name string
        num  uint32
        want int
    }{
        {"11 in binary", 0b1011, 3},
        {"128 in binary", 0b10000000, 1},
        {"2147483648", 2147483648, 1},
        {"zero", 0, 0},
        {"all ones 32-bit", 4294967295, 32},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := hammingWeight(tt.num); got != tt.want {
                t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
            }
        })
    }
}
