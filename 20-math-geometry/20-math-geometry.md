---
title: "Pattern: Math & Geometry"
aliases: [Pattern: Math & Geometry]
type: pattern
tags:
  - pattern
---

# Math & Geometry

## 🎯 The Core Idea

Math in coding isn't about memorizing formulas — it's about **number patterns**, **properties**, and **spatial reasoning**. Many problems reduce to clever observations.

**> Quick thought**: Why do math problems appear in coding interviews?

<details>
<summary>Click to reveal...</summary>

They test **pattern recognition** and **deep understanding** of how numbers work. Plus, many real-world problems (graphics, physics, finance) rely on math.
</details>

---

## 🔥 Essential Math Concepts

### 1. GCD and LCM

```go
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

// GCD of array
func findGCD(nums []int) int {
    result := nums[0]
    for _, num := range nums[1:] {
        result = gcd(result, num)
    }
    return result
}
```

**Euclidean algorithm**: Keep replacing (a, b) with (b, a mod b) until b=0.

---

### 2. Prime Numbers

```go
func isPrime(n int) bool {
    if n < 2 { return false }
    if n == 2 { return true }
    if n % 2 == 0 { return false }
    
    // Check up to sqrt(n)
    for i := 3; i*i <= n; i += 2 {
        if n % i == 0 { return false }
    }
    return true
}

// Sieve of Eratosthenes: All primes up to n
func sieve(n int) []bool {
    isPrime := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        isPrime[i] = true
    }
    
    for i := 2; i*i <= n; i++ {
        if isPrime[i] {
            for j := i * i; j <= n; j += i {
                isPrime[j] = false
            }
        }
    }
    return isPrime
}
```

**Key insight**: If n is composite, it has a factor ≤ √n.

---

### 3. Power with Modulo

```go
func pow(base, exp, mod int) int {
    result := 1
    base %= mod
    
    for exp > 0 {
        if exp & 1 == 1 {  // If exp is odd
            result = (result * base) % mod
        }
        base = (base * base) % mod
        exp >>= 1  // Divide by 2
    }
    return result
}
```

**Fast exponentiation**: O(log n) instead of O(n).

**Example**: 3^13 = 3 × 3^12 = 3 × (3^6)^2 = 3 × ((3^3)^2)^2

---

### 4. Number Theory Tricks

```go
// Count digits
func countDigits(n int) int {
    if n == 0 { return 1 }
    count := 0
    for n != 0 {
        count++
        n /= 10
    }
    return count
}

// Sum of digits
func sumDigits(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

// Reverse number
func reverse(n int) int {
    rev := 0
    for n != 0 {
        rev = rev*10 + n%10
        n /= 10
    }
    return rev
}

// Check power of 2
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n-1)) == 0
}
```

---

## 🔥 Geometry Concepts

### 1. Points, Lines, Slopes

```go
type Point struct {
    x, y int
}

// GCD of two numbers
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Unique lines through points
func maxPoints(points []Point) int {
    if len(points) <= 2 {
        return len(points)
    }
    
    maxCount := 0
    for i := 0; i < len(points); i++ {
        slopes := make(map[string]int)
        same := 1
        
        for j := i + 1; j < len(points); j++ {
            dx := points[j].x - points[i].x
            dy := points[j].y - points[i].y
            
            if dx == 0 && dy == 0 {
                same++
                continue
            }
            
            // Reduce slope to simplest form
            g := gcd(dx, dy)
            dx /= g
            dy /= g
            
            key := fmt.Sprintf("%d,%d", dx, dy)
            slopes[key]++
        }
        
        localMax := same
        for _, count := range slopes {
            localMax = max(localMax, count+same)
        }
        maxCount = max(maxCount, localMax)
    }
    
    return maxCount
}
```

**Key insight**: Represent slope as reduced fraction `dy/dx` to avoid floating point issues.

---

### 2. Rectangle Overlap

```go
func isRectangleOverlap(rec1, rec2 []int) bool {
    // Check if one rectangle is to the left
    if rec1[2] <= rec2[0] || rec2[2] <= rec1[0] {
        return false
    }
    // Check if one rectangle is above
    if rec1[3] <= rec2[1] || rec2[3] <= rec1[1] {
        return false
    }
    return true
}
```

**Format**: `[x1, y1, x2, y2]` where (x1, y1) is bottom-left, (x2, y2) is top-right.

---

## 🔥 Special Number Sequences

### Happy Number (Floyd's Cycle Detection)

```go
func isHappy(n int) bool {
    slow, fast := n, n
    
    for {
        slow = sumSquares(slow)
        fast = sumSquares(sumSquares(fast))
        
        if slow == fast {
            break
        }
    }
    
    return slow == 1
}

func sumSquares(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}
```

**Why it works**: The sequence either reaches 1 (happy) or enters a cycle (unhappy). Same as cycle detection in linked lists!

---

### Ugly Numbers

```go
func nthUglyNumber(n int) int {
    ugly := make([]int, n)
    ugly[0] = 1
    
    i2, i3, i5 := 0, 0, 0
    
    for i := 1; i < n; i++ {
        next2, next3, next5 := ugly[i2]*2, ugly[i3]*3, ugly[i5]*5
        ugly[i] = min(next2, next3, next5)
        
        if ugly[i] == next2 { i2++ }
        if ugly[i] == next3 { i3++ }
        if ugly[i] == next5 { i5++ }
    }
    
    return ugly[n-1]
}
```

**Key insight**: Each ugly number is a previous ugly number × 2, × 3, or × 5.

---

## 🎮 Practice Exercise

**> Problem: Given two integers a and b, return the sum of the series: a + aa + aaa + aaaa + ... (b terms)**

**Example**: a=2, b=3 → 2 + 22 + 222 = 246

<details>
<summary>Hint: How to build 2, 22, 222?</summary>

Each term = previous term × 10 + a
</details>

<details>
<summary>Solution</summary>

```go
func seriesSum(a, b int) int {
    term := 0
    sum := 0
    
    for i := 0; i < b; i++ {
        term = term*10 + a
        sum += term
    }
    
    return sum
}
```
</details>

---

## 📊 Common Formulas

| Concept | Formula |
|---------|---------|
| Sum of 1 to n | n(n+1)/2 |
| Sum of squares | n(n+1)(2n+1)/6 |
| Sum of cubes | [n(n+1)/2]² |
| n! mod p | Compute iteratively |
| nCr (combinations) | n!/(r! × (n-r)!) |

---

## ⚠️ Common Pitfalls

1. **Integer overflow**: Use `int64` for large calculations
2. **Division by zero**: Always check before dividing
3. **Floating point**: Use fractions instead of floats for precision
4. **Negative numbers**: GCD, modulo work differently with negatives
5. **Edge cases**: 0, 1, negative inputs

---

## 🚀 Problem Types

| Type | Pattern | Example |
|------|---------|---------|
| Number theory | GCD, primes, factors | Count primes, is prime |
| Geometry | Points, lines, overlap | Max points on line |
| Sequences | Find pattern, generate | Ugly numbers |
| Arithmetic | String operations | Add strings, multiply strings |
| Probability | Counting, combinations | Pascal's triangle |

---

## 💡 Pro Tips

1. **Test with examples**: Verify formulas with small numbers
2. **Handle negatives**: Make sure to account for negative inputs
3. **Use math tricks**: Properties like (a+b)² = a² + 2ab + b²
4. **Modulo arithmetic**: (a + b) mod m = [(a mod m) + (b mod m)] mod m
5. **Bit manipulation**: Often faster than arithmetic

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[23-matrix|Matrix]] — Matrix operations and geometry
- [[17-bit-manipulation|Bit Manipulation]] — Some math problems use bit tricks
