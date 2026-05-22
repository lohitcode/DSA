package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
    tests := []struct {
        s    string
        want int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3},
    }
    for _, tt := range tests {
        if got := lengthOfLongestSubstring(tt.s); got != tt.want {
            t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
        }
    }
}
