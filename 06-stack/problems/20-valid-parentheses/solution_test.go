package main

import "testing"

func TestIsValid(t *testing.T) {
    tests := []struct{ s string; want bool }{
        {"()", true},
        {"()[]{}", true},
        {"(]", false},
    }
    for _, tt := range tests {
        if got := isValid(tt.s); got != tt.want {
            t.Errorf("isValid() = %v, want %v", got, tt.want)
        }
    }
}
