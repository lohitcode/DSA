---
title: "Pattern: Bit Manipulation"
aliases: [Pattern: Bit Manipulation]
type: pattern
tags:
  - pattern
---

# Bit Manipulation

## 🎯 The Core Idea

Computers store everything in **binary** (0s and 1s). Bit manipulation works directly with these bits — incredibly fast and can solve problems that seem impossible otherwise.

```
Number:  5
Binary:  0 1 0 1
Bits:    8 4 2 1  (place values)
         └ └ └ └
         └─┴─┴─┴── 4 + 1 = 5
```

**> Quick thought**: Why bother with bits when high-level languages work fine?

<details>
<summary>Click to reveal...</summary>

Bit operations are O(1) and can solve certain problems elegantly (XOR tricks, set operations, space optimization). Also essential for systems programming, compression, and hashing.
</details>

---

## 🧠 Bit Operations (The Alphabet)

```go
a & b    // AND: 1 only if both bits are 1
a | b    // OR:  1 if either bit is 1
a ^ b    // XOR: 1 if bits are different
^a       // NOT: Flip all bits (0→1, 1→0)
a << n   // Left shift: Multiply by 2^n
a >> n   // Right shift: Divide by 2^n (floor)
```

### Visual Examples
```
AND:     5 & 3 = 1           OR:      5 | 3 = 7
  0101                       0101
& 0011                     | 0011
  ----                       ----
  0001 = 1                   0111 = 7

XOR:     5 ^ 3 = 6           Left:    5 << 2 = 20
  0101                       0101 × 4
^ 0011                     ---------
  ----                       10100 = 20
  0110 = 6
```

---

## 🔥 Essential Bit Tricks

### 1. Check if Odd/Even
```go
if n & 1 == 0 { /* even */ }  // Last bit is 0
if n & 1 == 1 { /* odd */ }   // Last bit is 1
```

### 2. Check if Power of 2
```go
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n-1)) == 0
}
```

**Why it works**: Powers of 2 have exactly ONE bit set.
```
8:  1000      8-1:  0111
16: 10000     16-1: 01111
AND = 0!
```

### 3. Get the Rightmost Set Bit
```go
rightmost := n & -n  // Isolates lowest set bit
```

**Example**: `12 & -12`
```
12:  1100
-12: 0100 (two's complement)
AND: 0100 = 4
```

### 4. Clear Rightmost Set Bit
```go
n &= n - 1  // Turns off the rightmost 1-bit
```

### 5. Count Set Bits (Brian Kernighan's Algorithm)
```go
func countBits(n int) int {
    count := 0
    for n != 0 {
        n &= n - 1  // Clear rightmost set bit
        count++
    }
    return count
}
```

**Time**: O(k) where k = number of set bits (not total bits!)

---

## 🔥 Common Bit Patterns

### Pattern 1: Find the Single Number
**> Every number appears twice except one. Find it.**

```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num  // XOR all numbers
    }
    return result
}
```

**Why it works**: `a ^ a = 0` and `a ^ 0 = a`. Pairs cancel out!

```
[2, 1, 4, 5, 2, 4, 1]
XOR all: 2^2 ^ 1^1 ^ 4^4 ^ 5 = 0 ^ 0 ^ 0 ^ 5 = 5
```

---

### Pattern 2: Find Two Unique Numbers
**> Every number appears twice except two. Find them.**

```go
func singleNumbers(nums []int) []int {
    // XOR all: result = a ^ b (the two unique numbers)
    xor := 0
    for _, num := range nums {
        xor ^= num
    }
    
    // Find rightmost different bit
    diff := xor & -xor
    
    // Separate numbers based on this bit
    a, b := 0, 0
    for _, num := range nums {
        if num&diff == 0 {
            a ^= num
        } else {
            b ^= num
        }
    }
    
    return []int{a, b}
}
```

**Key insight**: Any set bit where `a` and `b` differ can separate them into two groups.

---

### Pattern 3: Add Without +
**> Add two integers using bitwise ops**

```go
func getSum(a, b int) int {
    for b != 0 {
        carry := a & b      // Bits that will carry
        a = a ^ b           // Add without carry
        b = carry << 1      // Shift carry to next position
    }
    return a
}
```

**Visual**:
```
a = 5 (0101), b = 3 (0011)

Step 1: carry = 0001, a = 0110, b = 0010
Step 2: carry = 0010, a = 0100, b = 0100
Step 3: carry = 0100, a = 0000, b = 1000
Step 4: carry = 0000, a = 1000, b = 0000 → return 8
```

---

### Pattern 4: Swap Without Temp
```go
a, b = b, a  // In Go, this works!

// But the bit manipulation way:
a = a ^ b
b = a ^ b  // (a ^ b) ^ b = a
a = a ^ b  // (a ^ b) ^ a = b
```

---

### Pattern 5: Generate All Subsets
**> Use bitmask to represent subset**

```go
func subsets(nums []int) [][]int {
    n := len(nums)
    result := [][]int{}
    
    // 2^n subsets
    for mask := 0; mask < (1 << n); mask++ {
        subset := []int{}
        for i := 0; i < n; i++ {
            if mask&(1<<i) != 0 {
                subset = append(subset, nums[i])
            }
        }
        result = append(result, subset)
    }
    
    return result
}
```

**Visual** for `[1, 2, 3]`:
```
Mask  Binary  Subset
000   000     []
001   001     [1]
010   010     [2]
011   011     [1,2]
...
111   111     [1,2,3]
```

---

## 🎮 Practice Exercise

**> Problem**: Find the missing number in range [0, n] from an array of n distinct numbers.

<details>
<summary>Hint: XOR trick...</summary>

XOR all indices and all values. The missing number will remain!
</details>

<details>
<summary>Solution</summary>

```go
func missingNumber(nums []int) int {
    n := len(nums)
    result := n  // Start with n (since range is [0, n])
    
    for i := 0; i < n; i++ {
        result ^= i ^ nums[i]
    }
    
    return result
}
```
</details>

---

## 📊 Complexity

| Operation | Time | Notes |
|-----------|------|-------|
| All bit ops | O(1) | Constant time |
| Count bits | O(k) | k = set bits |
| Generate subsets | O(2^n × n) | Exponential |

---

## ⚠️ Common Pitfalls

1. **Precedence**: `&` has lower precedence than `==` — use parentheses!
2. **Signed vs unsigned**: Right shift on signed numbers can be arithmetic (fills with sign bit)
3. **Integer overflow**: Shifting can overflow — be careful with large values
4. **Endianness**: When dealing with byte arrays, consider endianness
5. **NOT on signed**: `^x` on signed int gives negative result

---

## 🚀 When to Use Bit Manipulation

✅ **Use Bit Manipulation when:**
- Need to find unique elements (XOR tricks)
- Working with sets/subsets (bitmask)
- Space optimization (pack multiple values in one int)
- Checking divisibility by powers of 2
- Fast multiplication/division by powers of 2

❌ **Don't use when:**
- Code readability is more important
- Problem doesn't involve binary properties
- Using high-level abstractions

---

## 💡 Real-World Applications

- **Hash functions**: Use XOR for combining hashes
- **Compression**: Huffman coding uses bit patterns
- **Cryptography**: AES, RSA work at bit level
- **Graphics**: Pixel manipulation, color channels
- **Network**: IP addresses, subnet masks
- **Flags**: Combine multiple boolean flags in one integer

---

## 🧩 Bitwise Cheat Sheet

```go
// Set bit at position i
n |= (1 << i)

// Clear bit at position i
n &= ^(1 << i)

// Toggle bit at position i
n ^= (1 << i)

// Check bit at position i
if n&(1<<i) != 0 { /* bit is set */ }

// Get lowest set bit
n & -n

// Clear lowest set bit
n & (n - 1)

// Check if power of 2
n > 0 && (n & (n - 1)) == 0
```

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[03-arrays-hashing|Arrays & Hashing]] — Bit operations for hash functions
- [[01-linked-list|Linked List]] — Bit manipulation for cycle detection
- [[05-binary-search|Binary Search]] — Some bit problems use binary search
