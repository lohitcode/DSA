---
title: "Pattern: Two Pointers"
aliases: [Pattern: Two Pointers]
type: pattern
tags:
  - pattern
---

# Two Pointers

## Concept
Use two pointers to traverse data structure simultaneously, often moving toward each other or at different speeds. Reduces O(n²) to O(n) by avoiding nested loops.

## Visual Intuition
```
Opposite Direction:          Same Direction:
[1, 2, 3, 4, 5]              [1, 2, 3, 4, 5]
 ↑           ↑               ↑    ↑
L           R               S    F

Find pair with sum X         Detect cycle / Find middle
```

## Internal Mechanics
The key insight: **Process elements in a single pass** by using the structure's properties (sorted order, linked list links, etc.) to eliminate unnecessary comparisons.

## Common Patterns

### 1. Opposite Direction (Sorted Array)
**When**: Find pair/triplet with sum in sorted array.

**Key insight**: In sorted array, moving left increases sum, moving right decreases sum.

```go
left, right := 0, len(nums)-1
for left < right {
    sum := nums[left] + nums[right]
    if sum == target {
        return []int{left, right}
    } else if sum < target {
        left++   // Need larger sum
    } else {
        right--  // Need smaller sum
    }
}
```

**Example**: Two Sum II (sorted input)
- Time: O(n) — one pass
- Space: O(1) — no extra data structure

### 2. Same Direction (Fast/Slow Pointers)
**When**: Detect cycles, find middle, remove Nth from end.

**Key insight**: Fast pointer covers "future" ground for slow pointer.

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next      // Move 1 step
    fast = fast.Next.Next // Move 2 steps
}
// slow is now at middle (or cycle exists if slow == fast)
```

**Applications**:
- **Middle element**: Fast reaches end, slow at middle
- **Cycle detection**: If cycle exists, fast catches slow
- **Nth from end**: Fast moves N steps first, then both move together

### 3. Three Pointers (for Triplets)
**When**: Find triplets with sum in sorted array.

```go
for i := 0; i < len(nums)-2; i++ {
    left, right := i+1, len(nums)-1
    for left < right {
        sum := nums[i] + nums[left] + nums[right]
        if sum == target {
            return triplet
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
}
```

### 4. Pointer Chasing (Linked Lists)
**When**: Merge two sorted lists, find intersection.

```go
for p1 != nil && p2 != nil {
    if p1.Val < p2.Val {
        result = p1
        p1 = p1.Next
    } else {
        result = p2
        p2 = p2.Next
    }
}
```

## When to Use

| Pattern | Condition | Example |
|---------|-----------|---------|
| Opposite pointers | Sorted array | Two Sum, Container with Water |
| Fast/slow | Linked list, cycle | Middle node, Cycle detect |
| Left/right on string | Palindrome, valid substring | Valid Palindrome |
| Merge pointers | Two sorted inputs | Merge Sorted Arrays |

## Complexity Analysis
- **Time**: O(n) — single pass (or O(n²) for triplets)
- **Space**: O(1) — only pointer variables

## Common Pitfalls
1. **Off-by-one**: `left < right` vs `left <= right` — depends on problem
2. **Skipping duplicates**: For triplet problems, skip same values to avoid duplicates
3. **Null checks**: Always check `fast != nil && fast.Next != nil`
4. **Infinite loop**: Ensure at least one pointer moves each iteration
5. **Overflow**: Sum of two numbers might exceed int range

## Optimization Scenarios

### From Brute Force to Two Pointers
```go
// Brute Force: O(n²)
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
        if nums[i] + nums[j] == target { return }
    }
}

// Two Pointers: O(n)
left, right := 0, n-1
for left < right {
    if nums[left] + nums[right] == target { return }
    // Move pointers
}
```

**Why it works**: Sorting allows us to make "informed" decisions about which direction to move.

## Related Patterns
- **Sliding Window**: Often uses two pointers for window boundaries
- **Linked List**: Fast/slow pointers fundamental to list problems
- **Binary Search**: Can be seen as shrinking range with two pointers

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[04-sliding-window|Sliding Window]] — Uses two pointers for window boundaries
- [[07-tree-dfs|Tree DFS]] — Fast/slow pointers for finding middle node
- [[22-string-advanced|String Advanced]] — Two pointers for palindrome checking
- [[01-linked-list|Linked List]] — Fast/slow pointer technique
