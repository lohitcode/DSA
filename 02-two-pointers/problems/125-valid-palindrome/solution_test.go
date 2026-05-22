package main

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct{ s string; want bool }{
        {"A man, a plan, a canal: Panama", true},
        {"race a car", false},
        {" ", true},
        {"a", true},
    }
    for _, tt := range tests {
        if got := isPalindrome(tt.s); got != tt.want {
            t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
        }
    }
}
