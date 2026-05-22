---
title: "Pattern: Stack"
aliases: [Pattern: Stack]
type: pattern
tags:
  - pattern
---

# Stack

## 🎯 The Core Idea

A stack is like a stack of plates: **you can only take from the top**. Last plate you put on is the first one you take off.

```
    ┌───┐
    │ 3 │ ← Top (most recent)
    ├───┤
    │ 2 │
    ├───┤
    │ 1 │ ← Bottom (oldest)
    └───┘
```

**LIFO**: Last In, First Out

**> Quick thought**: Why would you want to only access the top?

<details>
<summary>Click to reveal...</summary>

Some problems have **nested structure** — you process things in reverse order of how you saw them. Like undo/redo, browser back button, or function calls!
</details>

---

## 🧠 Stack Operations

```go
// Go uses slices as stacks
stack := []int{}

// Push: Add to top
stack = append(stack, 5)

// Pop: Remove from top
val := stack[len(stack)-1]
stack = stack[:len(stack)-1]

// Peek: Look at top without removing
top := stack[len(stack)-1]

// Empty check
len(stack) == 0
```

---

## 🔥 Common Stack Patterns

### Pattern 1: Parentheses Matching
**> Are the brackets balanced?**

```go
func isValid(s string) bool {
    stack := []rune{}
    matching := map[rune]rune{
        ')': '(', '}': '{', ']': '[',
    }
    
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)  // Push opening
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != matching[ch] {
                return false  // Mismatch!
            }
            stack = stack[:len(stack)-1]  // Pop matching
        }
    }
    
    return len(stack) == 0  // All matched?
}
```

**Key insight**: Stack remembers what you need to match later!

---

### Pattern 2: Monotonic Stack
**> Keep stack in sorted order while processing**

```go
// Find next greater element
func nextGreaterElement(nums []int) []int {
    result := make([]int, len(nums))
    for i := range result {
        result[i] = -1  // Default: no greater element
    }
    
    stack := []int{}  // Indices of elements
    
    for i := range nums {
        for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
            // Current is greater than stack's top
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[top] = nums[i]
        }
        stack = append(stack, i)
    }
    
    return result
}
```

**What's happening**: Stack holds elements waiting for a "greater" neighbor. When we find one, we resolve all waiting elements smaller than current.

---

### Pattern 3: Removing Duplicates
**> Remove adjacent duplicates repeatedly**

```go
func removeDuplicates(s string) string {
    stack := []rune{}
    
    for _, ch := range s {
        if len(stack) > 0 && stack[len(stack)-1] == ch {
            stack = stack[:len(stack)-1]  // Pop duplicate
        } else {
            stack = append(stack, ch)  // Push
        }
    }
    
    return string(stack)
}
```

**Example**: `azxxzy` → `ay` (xx removed, then zz removed)

---

### Pattern 4: Daily Temperatures
**> How many days until warmer temperature?**

```go
func dailyTemperatures(temps []int) []int {
    result := make([]int, len(temps))
    stack := []int{}  // Indices of days waiting for warmer temp
    
    for i, temp := range temps {
        for len(stack) > 0 && temp > temps[stack[len(stack)-1]] {
            prevDay := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[prevDay] = i - prevDay  // Days waited
        }
        stack = append(stack, i)
    }
    
    return result  // Unresolved days stay 0
}
```

---

## 🎮 Practice Exercise

**> Problem**: Given a string, evaluate the expression. It can contain +, -, *, / and parentheses.

<details>
<summary>Hint: You might need TWO stacks...</summary>

One for numbers, one for operators. Process * and / immediately, but wait for ().
</details>

<details>
<summary>Solution approach</summary>

```go
func calculate(s string) int {
    nums, ops := []int{}, []byte{}
    i := 0
    
    for i < len(s) {
        ch := s[i]
        
        if ch == ' ' {
            i++
            continue
        }
        
        if isDigit(ch) {
            // Parse full number
            num := 0
            for i < len(s) && isDigit(s[i]) {
                num = num*10 + int(s[i]-'0')
                i++
            }
            nums = append(nums, num)
            continue
        }
        
        if ch == '(' {
            ops = append(ops, ch)
        } else if ch == ')' {
            // Solve until matching '('
            for ops[len(ops)-1] != '(' {
                solve(nums, ops)
            }
            ops = ops[:len(ops)-1]  // Pop '('
        } else {
            // Current op has lower/equal precedence? Solve first
            for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
                solve(nums, ops)
            }
            ops = append(ops, ch)
        }
        i++
    }
    
    for len(ops) > 0 {
        solve(nums, ops)
    }
    
    return nums[len(nums)-1]
}
```
</details>

---

## 📊 Complexity

| Operation | Time | Notes |
|-----------|------|-------|
| Push | O(1) | Amortized |
| Pop | O(1) | |
| Peek | O(1) | |
| Search | O(n) | Must scan all elements |

**Space**: O(n) for stack storage

---

## ⚠️ Common Pitfalls

1. **Empty stack access**: Always check `len(stack) > 0` before peek/pop
2. **Wrong order**: Remember LIFO — last in, first out!
3. **Off-by-one**: `stack[:len(stack)-1]` vs `stack[:len(stack)]`
4. **Forgotten elements**: Check if stack is empty at end
5. **Monotonic confusion**: Increasing vs decreasing stack — trace through example

---

## 🚀 When to Use a Stack

✅ **Use Stack when:**
- Matching pairs (parentheses, HTML tags)
- Reversal needed (palindrome, undo operation)
- Nested structures (expression evaluation, DFS)
- Tracking "previous" elements (monotonic stack)
- Removing adjacent duplicates

❌ **Don't use Stack when:**
- Need random access (use array)
- Need FIFO order (use queue)
- Need to search elements (use set/map)

---

## 💡 Related Patterns

- **Queue**: FIFO — use for BFS, level-order traversal
- **Monotonic Queue**: Similar idea, but pop from front too
- **Call Stack**: Recursion uses stack implicitly

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[01-linked-list|Linked List]] — Stacks can be implemented with linked lists
- [[07-tree-dfs|Tree DFS]] — Recursion uses call stack, explicit stack for iterative
- [[09-tree-bfs|Tree BFS]] — Stack (LIFO) vs Queue (FIFO) comparison
- [[21-design/21-design-problems|Design Problems]] — Min Stack design pattern
- [[11-backtracking|Backtracking]] — Uses stack for state management
