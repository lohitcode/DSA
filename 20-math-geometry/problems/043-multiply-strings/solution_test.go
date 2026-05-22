package main

import "testing"

func TestMultiply(t *testing.T) {
    tests := []struct {
        num1, num2, want string
    }{
        {"2", "3", "6"},
        {"123", "456", "56088"},
        {"0", "123", "0"},
    }
    for _, tt := range tests {
        if got := multiply(tt.num1, tt.num2); got != tt.want {
            t.Errorf("multiply() = %v, want %v", got, tt.want)
        }
    }
}
