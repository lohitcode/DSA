package main

import "testing"

func TestReverseBits(t *testing.T) {
    tests := []struct {
        num  uint32
        want uint32
    }{
        {0b00000010100101000001111010011100, 964176192},
        {0b11111111111111111111111111111101, 3221225471},
    }
    for _, tt := range tests {
        if got := reverseBits(tt.num); got != tt.want {
            t.Errorf("reverseBits() = %v, want %v", got, tt.want)
        }
    }
}
