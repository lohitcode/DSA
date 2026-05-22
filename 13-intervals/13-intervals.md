---
title: "Pattern: Intervals"
aliases: [Pattern: Intervals]
type: pattern
tags:
  - pattern
---

# Intervals

## 🎯 The Core Idea

Intervals represent **ranges** — like meeting times `[start, end]`. The key insight: **sorting by start time** reveals overlaps and makes problems tractable.

```
Before:              After sorting by start:
[5, 8]    [2, 4]     [2, 4]───[3, 5]───[5, 8]───[9, 12]
   [9, 12]  [3, 5]
    
Unsorted!            Easy to see overlaps now
```

**> Quick thought**: Why does sorting by start time work so well?

<details>
<summary>Click to reveal...</summary>

After sorting by start, you know each interval can only overlap with what comes **after** it. You can process greedily — make one pass and be done.
</details>

---

## 🧠 The Golden Rule: Sort First!

```go
// Step 1: ALWAYS sort by start time
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i].Start < intervals[j].Start
})

// Step 2: Single pass through sorted intervals
```

**Why**: After sorting, each interval's "neighbors" are the only ones that can possibly overlap.

---

## 🔥 Common Interval Patterns

### Pattern 1: Merge Overlapping Intervals
**> Combine overlapping ranges into single intervals**

```go
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 { return [][]int{} }
    
    // Step 1: Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    result := [][]int{intervals[0]}
    
    for _, interval := range intervals[1:] {
        last := result[len(result)-1]
        
        // Overlap? Current start <= last end
        if interval[0] <= last[1] {
            // Merge: extend end if needed
            last[1] = max(last[1], interval[1])
        } else {
            // No overlap: add new interval
            result = append(result, interval)
        }
    }
    
    return result
}
```

**Example**: `[1,3], [2,6], [8,10]` → `[1,6], [8,10]`

**Key insight**: After sorting, just compare each interval with the **last merged** one.

---

### Pattern 2: Insert Interval
**> Add new interval to sorted, non-overlapping list**

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    i := 0
    n := len(intervals)
    
    // Phase 1: Add all intervals ending before newInterval
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }
    
    // Phase 2: Merge all overlapping intervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    result = append(result, newInterval)
    
    // Phase 3: Add remaining intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}
```

**Key insight**: Three phases — before, merge, after.

---

### Pattern 3: Meeting Rooms (Overlap Check)
**> Can one person attend all meetings?**

```go
func canAttendMeetings(intervals [][]int) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    for i := 1; i < len(intervals); i++ {
        // Overlap exists if current start < previous end
        if intervals[i][0] < intervals[i-1][1] {
            return false
        }
    }
    
    return true
}
```

**Variation**: Minimum meeting rooms needed?
```go
func minMeetingRooms(intervals [][]int) int {
    // Separate start and end times
    starts := make([]int, len(intervals))
    ends := make([]int, len(intervals))
    
    for i, iv := range intervals {
        starts[i] = iv[0]
        ends[i] = iv[1]
    }
    
    sort.Ints(starts)
    sort.Ints(ends)
    
    rooms := 0
    endPtr := 0
    
    for _, start := range starts {
        // If this meeting starts after the earliest one ends
        if start >= ends[endPtr] {
            endPtr++  // Free up a room
        } else {
            rooms++  // Need new room
        }
    }
    
    return rooms
}
```

**Key insight**: Track when meetings END — that's when rooms free up.

---

### Pattern 4: Interval List Intersections
**> Find common overlap points in two interval lists**

```go
func intervalIntersection(A, B [][]int) [][]int {
    result := [][]int{}
    i, j := 0, 0
    
    for i < len(A) && j < len(B) {
        // Check for overlap
        start := max(A[i][0], B[j][0])
        end := min(A[i][1], B[j][1])
        
        if start <= end {
            result = append(result, []int{start, end})
        }
        
        // Move the one that ends first
        if A[i][1] < B[j][1] {
            i++
        } else {
            j++
        }
    }
    
    return result
}
```

**Key insight**: Two-pointer technique — advance the one that ends earlier.

---

## 🎮 Practice Exercise

**> Problem**: Given a list of intervals, remove the minimum number to make the rest non-overlapping.

<details>
<summary>Hint: This is like interval scheduling...</summary>

Classic greedy: Always pick the interval that ends earliest. This leaves maximum room for others.
</details>

<details>
<summary>Solution</summary>

```go
func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 { return 0 }
    
    // Sort by END time (not start!)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    
    count := 1  // Keep first interval
    lastEnd := intervals[0][1]
    
    for _, interval := range intervals[1:] {
        if interval[0] >= lastEnd {
            // No overlap, keep it
            count++
            lastEnd = interval[1]
        }
        // Else: overlap, skip (this is what we remove)
    }
    
    return len(intervals) - count
}
```
</details>

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(n log n) | Sorting dominates |
| Space | O(1) or O(n) | O(1) if sorting in-place, O(n) for result |

---

## ⚠️ Common Pitfalls

1. **Forgot to sort**: Patterns break without sorting by start time
2. **Wrong overlap condition**: `interval[0] <= last[1]` NOT `<` (edge case: touching at a point)
3. **Sort by wrong key**: For "min removals", sort by END time
4. **Off-by-one**: Inclusive vs exclusive endpoints affects overlap logic
5. **Modifying input**: Make copy if you need original later

---

## 🚀 When to Use Interval Patterns

✅ **Use when:**
- Working with time ranges, schedules
- Detecting overlaps between ranges
- Merging or splitting ranges
- Finding free/busy slots

**Common applications**:
- Calendar apps (meeting scheduling)
- CPU scheduling
- Network bandwidth allocation
- Genome sequence alignment

---

## 💡 Interval Representation

```go
// Option 1: Slice of 2 elements
interval := []int{start, end}

// Option 2: Struct (more readable)
type Interval struct {
    Start, End int
}
```

**Tip**: Use structs for production code, slices for quick LeetCode.

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[03-arrays-hashing|Arrays & Hashing]] — Interval processing on sorted arrays
- [[14-greedy|Greedy]] — Interval scheduling uses greedy approach
- [[02-two-pointers|Two Pointers]] — Merging intervals uses two pointers
