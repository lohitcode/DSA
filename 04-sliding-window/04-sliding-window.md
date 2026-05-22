---
title: "Pattern: Sliding Window"
aliases: [Pattern: Sliding Window]
type: pattern
tags:
  - pattern
---

# Sliding Window

## Concept
A technique to process subarrays of fixed or variable size using two pointers. Instead of recomputing for each window, we **incrementally update** the window state as it slides — achieving O(n) instead of O(n×k).

## Visual Intuition
```
Fixed Window (k=3):
[1, 2, 3, 4, 5, 6]
 └───┘     └───┘
window 1   window 2

Variable Window (sum ≥ 7):
[1, 2, 3, 4, 5, 6]
 └──────┘  └──┘
shrink    expand
```

## Internal Mechanics
The key insight: **When the window slides, only the elements entering and leaving change**.

**Naive approach**: For each window, recompute everything → O(n × k)
**Sliding window**: Update incrementally → O(n)

```
New sum = Old sum - element_leaving + element_entering
```

## Common Patterns

### 1. Fixed Size Window
**Template**: Find condition in every window of size k.

```go
// Step 1: Build first window
for i := 0; i < k; i++ {
    // Add nums[i] to window state
}

// Step 2: Slide window
for right := k; right < len(nums); right++ {
    // Remove nums[right-k] (left edge)
    // Add nums[right] (right edge)
    // Check/update result
}
```

**Example**: Max sum of subarray of size k
```go
maxSum, windowSum := 0, 0
// Build first window
for i := 0; i < k; i++ {
    windowSum += nums[i]
}
maxSum = windowSum

// Slide
for right := k; right < len(nums); right++ {
    windowSum += nums[right] - nums[right-k]
    maxSum = max(maxSum, windowSum)
}
```

### 2. Variable Size Window (Expand/Shrink)
**Template**: Find smallest/largest window satisfying condition.

```go
left := 0
for right := 0; right < len(nums); right++ {
    // Expand: add nums[right] to window
    for conditionSatisfied() {
        // Update result
        // Shrink: remove nums[left] from window
        left++
    }
}
```

**Example**: Smallest subarray with sum ≥ target
```go
left, minLen := 0, math.MaxInt
windowSum := 0

for right := 0; right < len(nums); right++ {
    windowSum += nums[right]
    for windowSum >= target {
        minLen = min(minLen, right-left+1)
        windowSum -= nums[left]
        left++
    }
}
```

### 3. Two Pointers (Sliding Window variant)
**When**: Find pairs/triplets with condition.

```go
left, right := 0, len(nums)-1
for left < right {
    sum := nums[left] + nums[right]
    if sum == target {
        return []int{left, right}
    } else if sum < target {
        left++
    } else {
        right--
    }
}
```

## When to Use Each

| Pattern | When | Key Insight |
|---------|------|-------------|
| Fixed window | "Every k elements" | Window size is known |
| Variable window | "Smallest window with condition" | Shrink when valid, expand when invalid |
| Two pointers | "Find pairs in sorted array" | Move pointers based on comparison |

## Complexity Analysis
- **Time**: O(n) — each element added/removed at most once
- **Space**: O(1) or O(k) — window state (usually O(1))

## Common Pitfalls
1. **Window boundaries**: Off-by-one errors (`right-left+1` vs `right-left`)
2. **Shrinking condition**: `while` vs `for` — use `while` for proper shrinking
3. **Empty result**: Handle case where no valid window exists
4. **Negative numbers**: Variable window shrinks on invalid, but negatives complicate "invalid" definition
5. **Integer overflow**: Window sum can exceed int range for large inputs

## Window State Management
Instead of recomputing, maintain a running state:
```go
// Instead of: sum(window) each time
windowSum := 0

// Instead of: max(window) each time
windowMax := 0  // Use deque for O(1) max/min

// Instead of: countDistinct(window) each time
windowCount := make(map[int]int)
```

## Advanced: With Hash Map / Frequency Counter
```go
windowCount := make(map[byte]int)
left := 0

for right := 0; right < len(s); right++ {
    // Add to window
    windowCount[s[right]]++

    // Shrink while invalid
    for !isValid(windowCount) {
        windowCount[s[left]]--
        if windowCount[s[left]] == 0 {
            delete(windowCount, s[left])
        }
        left++
    }
    // Update result
}
```

## Related Patterns
- **Two Pointers**: Often uses sliding window technique
- **Monotonic Deque**: For O(1) max/min in sliding window
- **Prefix Sum**: Can optimize some sliding window problems

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[02-two-pointers|Two Pointers]] — Uses two pointers for window boundaries
- [[22-string-advanced|String Advanced]] — Sliding window for substring problems
- [[03-arrays-hashing|Arrays & Hashing]] — Hash map tracks window elements
- [[12-heap|Heap]] — Monotonic deque for O(1) max/min in window
