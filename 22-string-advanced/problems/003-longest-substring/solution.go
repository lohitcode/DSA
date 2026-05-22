package main

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    left, maxLen := 0, 0
    for right, ch := range s {
        if idx, ok := charIndex[ch]; ok && idx >= left {
            left = idx + 1
        }
        charIndex[ch] = right
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
