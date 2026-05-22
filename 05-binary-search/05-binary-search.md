---
title: "Pattern: Binary Search"
aliases: [Pattern: Binary Search]
type: pattern
tags:
  - pattern
---

# Binary Search

## 🎯 The Core Idea

You're looking for a word in a **dictionary**. Do you start from page 1 and read every word? No! You open to the middle, see if your word comes before or after, and repeat.

That's binary search: **eliminate half the search space each time**.

```
[1, 3, 5, 7, 9, 11, 13, 15]
     ↑ mid = 7
Target: 13

7 < 13, so search RIGHT half:
            [9, 11, 13, 15]
                 ↑ mid = 13 ✓
```

**> Quick thought**: Why must the array be sorted?

<details>
<summary>Click to reveal...</summary>

Binary search RELIES on order. If the array isn't sorted, we can't eliminate half — we don't know which half to search!
</details>

---

## 🧠 The Power of O(log n)

| n | Linear O(n) | Binary O(log n) |
|---|-------------|-----------------|
| 100 | 100 steps | 7 steps |
| 1,000,000 | 1M steps | 20 steps |
| 1B | 1B steps | 30 steps |

**Mind-blowing**: Searching 1 billion elements takes only ~30 comparisons!

---

## 📝 The Three Variants (Know All Three!)

### Variant 1: Standard Binary Search
**> Find EXACT match or return -1**

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2  // Prevent overflow
        
        if nums[mid] == target {
            return mid  // Found!
        } else if nums[mid] < target {
            left = mid + 1  // Search right half
        } else {
            right = mid - 1  // Search left half
        }
    }
    
    return -1  // Not found
}
```

**Key details**:
- `left <= right`: Important! Use `<=` for standard search
- `mid + 1` / `mid - 1`: Move past mid to avoid infinite loop
- `left + (right-left)/2`: Prevents integer overflow

---

### Variant 2: Lower Bound (First ≥ Target)
**> Find first element NOT LESS than target**

```go
func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)  // Note: right = len(nums), not len(nums)-1
    
    for left < right {  // Note: < not <=
        mid := left + (right-left)/2
        
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid  // Don't exclude mid
        }
    }
    
    return left  // Index of first ≥ target, or len(nums) if none
}
```

**When to use**: Finding insertion point, counting elements < target

---

### Variant 3: Upper Bound (First > Target)
**> Find first element GREATER than target**

```go
func upperBound(nums []int, target int) int {
    left, right := 0, len(nums)
    
    for left < right {
        mid := left + (right-left)/2
        
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left  // Index of first > target
}
```

**When to use**: Finding range of elements, counting elements ≤ target

---

## 🔥 Advanced Pattern: Binary Search on Answer

**> When the array isn't sorted, but the ANSWER is!**

Sometimes you don't search the array — you search the **space of possible answers**.

**Example**: Koko eating bananas
```go
// Koko can eat at most K bananas per hour. Find minimum K.

func minEatingSpeed(piles []int, h int) int {
    // Search space: [1, max(piles)]
    left, right := 1, max(piles...)
    
    for left < right {
        mid := left + (right-left)/2
        
        if canEatAll(piles, h, mid) {
            right = mid  // Try slower
        } else {
            left = mid + 1  // Need faster
        }
    }
    
    return left
}

func canEatAll(piles []int, h, k int) bool {
    hours := 0
    for _, pile := range piles {
        hours += (pile + k - 1) / k  // Ceiling division
    }
    return hours <= h
}
```

**Key insight**: The "can eat all" function is monotonic — if K works, any K' > K also works. This monotonicity allows binary search!

---

## 🎮 Practice Exercise

**> Problem**: Given a sorted array that was rotated, find the minimum element.

```
Example: [3, 4, 5, 1, 2]
Minimum is 1 at index 3
```

<details>
<summary>Hint: What property does the rotated array have?</summary>

One half is always sorted. Compare mid with the rightmost element.
</details>

<details>
<summary>Solution</summary>

```go
func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if nums[mid] > nums[right] {
            // Min must be in right half
            left = mid + 1
        } else {
            // Min is in left half (including mid)
            right = mid
        }
    }
    
    return nums[left]
}
```
</details>

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(log n) | Halve search space each iteration |
| Space | O(1) | Only a few variables |

---

## ⚠️ Common Pitfalls (Don't Make These!)

1. **Infinite loop**: Using `left <= right` when you should use `left < right`
2. **Wrong bounds**: `right = mid` vs `right = mid - 1` — trace through!
3. **Integer overflow**: `mid := (left + right) / 2` can overflow!
4. **Off-by-one**: `len(nums)` vs `len(nums) - 1` for upper bound
5. **Monotonicity check**: For "binary search on answer", verify the predicate is actually monotonic!

---

## 🚀 When to Use Binary Search

✅ **Use Binary Search when:**
- Array is sorted (or you can make it monotonic)
- Need O(log n) lookups
- Finding "first/last" occurrence
- Searching answer space (min/max problems)

❌ **Don't use Binary Search when:**
- Array is unsorted and unsortable
- Need to process all elements anyway
- Array is small (O(n) is simpler)

---

## 💡 Pro Tips

1. **Always trace**: Run through example with pen and paper
2. **Check bounds**: After loop, verify `left` and `right` values
3. **Use templates**: Memorize the 3 variants, don't reinvent each time
4. **Test edge cases**: Empty array, single element, target not found

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[19-bst/19-binary-search-tree|Binary Search Tree]] — BST property enables binary search on trees
- [[14-greedy|Greedy]] — Binary search on answer is a greedy approach
- [[08-dp-1d|DP 1D]] — Some DP problems can be optimized with binary search
