---
title: "Pattern: String Advanced"
aliases: [Pattern: String Advanced]
type: pattern
tags:
  - pattern
---

# String Advanced

## 🎯 The Core Idea

Strings are just **arrays of characters**, but they have special patterns and operations. Advanced string problems often combine multiple techniques: two pointers, sliding window, DP, or even treat strings as graphs.

**> Quick thought**: Why are strings so common in coding problems?

<details>
<summary>Click to reveal...</summary>

Strings represent text, DNA, paths — real data! Plus, they're perfect for testing multiple skills: iteration, manipulation, comparison, and sometimes even math.
</details>

---

## 🔥 Essential String Patterns

### Pattern 1: Two Pointers for Strings
**> Compare from both ends moving inward**

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    
    for left < right {
        // Skip non-alphanumeric
        for left < right && !isAlphaNum(s[left]) {
            left++
        }
        for left < right && !isAlphaNum(s[right]) {
            right--
        }
        
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }
        left++
        right--
    }
    return true
}
```

**Common uses**: Palindrome checking, reversing, removing duplicates

---

### Pattern 2: Sliding Window for Substrings
**> Find longest/shortest substring with property**

```go
func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    left, maxLen := 0, 0
    
    for right, ch := range s {
        // If char seen and in current window
        if idx, ok := charIndex[ch]; ok && idx >= left {
            left = idx + 1  // Shrink from left
        }
        
        charIndex[ch] = right
        maxLen = max(maxLen, right-left+1)
    }
    
    return maxLen
}
```

**Variations**:
- Longest substring with at most K distinct characters
- Smallest window containing all characters of pattern
- Longest substring with repeating characters

---

### Pattern 3: String Matching
**> Find pattern in text**

**Naive**: Check every position → O(n×m)
**KMP**: Use failure function → O(n+m)
**Rabin-Karp**: Use hashing → O(n+m) average

```go
// Rabin-Karp (simplified)
func strStr(haystack, needle string) int {
    if len(needle) == 0 { return 0 }
    if len(needle) > len(haystack) { return -1 }
    
    // Simple hash: sum of character codes
    needleHash := hash(needle)
    windowHash := hash(haystack[:len(needle)])
    
    if needleHash == windowHash && haystack[:len(needle)] == needle {
        return 0
    }
    
    for i := len(needle); i < len(haystack); i++ {
        // Rolling hash: remove outgoing, add incoming
        windowHash += int(haystack[i]) - int(haystack[i-len(needle)])
        
        if windowHash == needleHash && haystack[i-len(needle)+1:i+1] == needle {
            return i - len(needle) + 1
        }
    }
    
    return -1
}

func hash(s string) int {
    h := 0
    for _, ch := range s {
        h += int(ch)
    }
    return h
}
```

---

### Pattern 4: Anagram Groups
**> Group strings that are anagrams**

```go
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    
    for _, s := range strs {
        // Sort string to get key
        key := sortString(s)
        groups[key] = append(groups[key], s)
    }
    
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}

func sortString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}
```

**Optimization**: Use count array as key instead of sorting.

---

### Pattern 5: String Transformation
**> Apply operations to transform string**

**Valid Palindrome II**: Can you make palindrome by deleting at most one character?

```go
func validPalindrome(s string) bool {
    return isPal(s, 0, len(s)-1, false)
}

func isPal(s string, left, right int, deleted bool) bool {
    for left < right {
        if s[left] != s[right] {
            if deleted {
                return false  // Already used our one deletion
            }
            // Try deleting left OR right
            return isPal(s, left+1, right, true) ||
                   isPal(s, left, right-1, true)
        }
        left++
        right--
    }
    return true
}
```

---

## 🔥 String Math Operations

### Add Strings
```go
func addStrings(num1, num2 string) string {
    i, j := len(num1)-1, len(num2)-1
    carry := 0
    result := []byte{}
    
    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 {
            sum += int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(num2[j] - '0')
            j--
        }
        
        result = append(result, byte(sum%10)+'0')
        carry = sum / 10
    }
    
    // Reverse result
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result)
}
```

---

## 🎮 Practice Exercise

**> Problem: Given two strings, check if they're one edit away (insert, delete, or replace one character).**

<details>
<summary>Think about the three cases...</summary>

- Same length: Check for one replacement
- Different lengths: Check if inserting one char makes them equal
</details>

<details>
<summary>Solution</summary>

```go
func oneEditAway(first, second string) bool {
    m, n := len(first), len(second)
    
    if abs(m - n) > 1 {
        return false
    }
    
    i, j := 0, 0
    foundDifference := false
    
    for i < m && j < n {
        if first[i] != second[j] {
            if foundDifference {
                return false  // Already found a difference
            }
            foundDifference = true
            
            if m > n {
                i++  // Delete from first (or insert in second)
            } else if m < n {
                j++  // Delete from second (or insert in first)
            } else {
                i++  // Replacement, move both
                j++
            }
        } else {
            i++
            j++
        }
    }
    
    return true
}
```
</details>

---

## 📊 String Operation Complexity

| Operation | Time | Notes |
|-----------|------|-------|
| Concatenation (+) | O(n+m) | Creates new string |
| strings.Builder | O(1) amortized | Efficient for many appends |
| Substring | O(1) | In Go, shares underlying array |
| Split | O(n) | Scan entire string |
| Replace | O(n) | Single pass |
| Contains | O(n×m) | Naive, O(n+m) with KMP |

---

## ⚠️ Common Pitfalls

1. **String concatenation in loop**: Use `strings.Builder`
2. **Rune vs byte**: Use range for Unicode, index for ASCII
3. **Immutable**: Strings can't be modified, create new ones
4. **Off-by-one**: Indexing is 0-based, length is actual count
5. **Empty strings**: Always handle `""` case

---

## 🚀 When to Use Each Pattern

| Pattern | Use when... |
|---------|------------|
| Two pointers | Comparing ends, reversing |
| Sliding window | Longest/shortest substring |
| Hash map | Counting characters, anagrams |
| Two strings | Compare, merge, transform |
| DP | Subsequence, edit distance |

---

## 💡 Go String Tips

```go
// Efficient building
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("hello")
}
result := builder.String()

// Character iteration
for i, ch := range s {  // ch is rune (Unicode)
    fmt.Printf("%d: %c\n", i, ch)
}

// Byte iteration (ASCII only)
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %c\n", i, s[i])
}

// String to rune slice
runes := []rune(s)
sort.Slice(runes, func(i, j int) bool {
    return runes[i] < runes[j]
})

// Rune slice to string
s = string(runes)
```

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[04-sliding-window|Sliding Window]] — For substring problems
- [[02-two-pointers|Two Pointers]] — For palindrome checking
- [[15-trie|Trie]] — For prefix-based string operations
