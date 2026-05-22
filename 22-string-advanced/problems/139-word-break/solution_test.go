package main

import "testing"

func TestWordBreak(t *testing.T) {
    tests := []struct {
        s        string
        wordDict []string
        want     bool
    }{
        {"leetcode", []string{"leet", "code"}, true},
        {"applepenapple", []string{"apple", "pen"}, true},
        {"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
    }
    for _, tt := range tests {
        if got := wordBreak(tt.s, tt.wordDict); got != tt.want {
            t.Errorf("wordBreak() = %v, want %v", got, tt.want)
        }
    }
}
