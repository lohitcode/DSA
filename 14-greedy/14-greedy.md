---
title: "Pattern: Greedy"
aliases: [Pattern: Greedy]
type: pattern
tags:
  - pattern
---

# Greedy Algorithms

## 🎯 The Core Idea

Greedy = always make the **locally optimal** choice. Hope it leads to **global optimum**.

```
At each step, pick what seems best RIGHT NOW.
Don't look ahead. Don't reconsider.
```

**> Quick thought**: When does greedy actually work?

<details>
<summary>Click to reveal...</summary>

Greedy works when the problem has **optimal substructure** AND **greedy choice property**. Translation: local optimal choices lead to global optimum. Not all problems have this!
</details>

---

## 🧠 Greedy vs. Dynamic Programming

| Greedy | Dynamic Programming |
|--------|-------------------|
| Make one choice, never revisit | Explore all choices, pick best |
| Fast, simple | Slower, more complex |
| Works for special problems | Works for many problems |
| Example: Dijkstra (shortest path) | Example: Floyd-Warshall (all pairs) |

**Key question**: Can I prove that making the greedy choice now is never wrong?

---

## 🔥 Common Greedy Patterns

### Pattern 1: Jump Game
**> Can I reach the end?**

```go
func canJump(nums []int) bool {
    maxReach := 0
    
    for i := 0; i < len(nums); i++ {
        // If I can't reach this position, I'm stuck
        if i > maxReach {
            return false
        }
        // Update farthest I can reach
        maxReach = max(maxReach, i + nums[i])
        
        // Early exit
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    
    return maxReach >= len(nums)-1
}
```

**Greedy insight**: Track the **farthest reachable position**. If at any point current position > maxReach, it's impossible.

---

### Pattern 2: Interval Scheduling
**> Maximum number of non-overlapping intervals**

```go
func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 { return 0 }
    
    // Greedy: Pick earliest finishing interval first
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    
    count := 1
    lastEnd := intervals[0][1]
    
    for _, interval := range intervals[1:] {
        if interval[0] >= lastEnd {
            count++
            lastEnd = interval[1]
        }
    }
    
    return len(intervals) - count  // Intervals to remove
}
```

**Greedy insight**: Picking the interval that **ends earliest** leaves maximum room for future intervals.

**Why it works**: If you have an optimal solution, you can always replace the first interval with the earliest-finishing one without making the solution worse.

---

### Pattern 3: Coin Change (Special Case)
**> Minimum coins when denominations are "canonical"**

```go
// ONLY works for certain coin systems like US coins
// Does NOT work for arbitrary coin systems!
func coinChange(coins []int, amount int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(coins)))
    
    count := 0
    for _, coin := range coins {
        count += amount / coin
        amount %= coin
    }
    
    if amount != 0 { return -1 }
    return count
}
```

**Warning**: This greedy approach FAILS for coins like `{1, 3, 4}` trying to make 6:
- Greedy: 4 + 1 + 1 = 3 coins
- Optimal: 3 + 3 = 2 coins

**Real solution**: Use DP for arbitrary coin systems.

---

### Pattern 4: Gas Station
**> Can you travel around the circuit?**

```go
func canCompleteCircuit(gas, cost []int) int {
    totalTank, currTank, start := 0, 0, 0
    
    for i := 0; i < len(gas); i++ {
        totalTank += gas[i] - cost[i]
        currTank += gas[i] - cost[i]
        
        // If we can't reach i+1 from start
        if currTank < 0 {
            // Start from i+1 instead
            start = i + 1
            currTank = 0
        }
    }
    
    if totalTank >= 0 {
        return start
    }
    return -1
}
```

**Greedy insight**: If you can't reach station B from station A, then no station between A and B can be the starting point either.

---

### Pattern 5: Partition Labels
**> Partition string into as many parts as possible**

```go
func partitionLabels(s string) []int {
    // Last occurrence of each character
    last := make(map[rune]int)
    for i, ch := range s {
        last[ch] = i
    }
    
    result := []int{}
    start, end := 0, 0
    
    for i, ch := range s {
        end = max(end, last[ch])  // Extend partition
        
        // Reached the end of current partition
        if i == end {
            result = append(result, end - start + 1)
            start = i + 1
        }
    }
    
    return result
}
```

**Greedy insight**: Extend each partition as far as necessary (to the last occurrence of any character in it), then cut.

---

## 🎮 Practice Exercise

**> Problem**: You're given tasks with deadlines and profits. Do one task at a time. Maximize profit.

<details>
<summary>Hint: What should you prioritize?</summary>

Highest profit tasks first! But make sure they can be scheduled before their deadline.
</details>

<details>
<summary>Approach</summary>

1. Sort tasks by profit (descending)
2. For each task, schedule it as late as possible before its deadline
3. Use union-find or simple array to find available slots

```go
// Pseudocode
sort tasks by profit (high to low)
slots = array of size maxDeadline

for task in tasks:
    // Find latest available slot before deadline
    for slot = task.deadline - 1; slot >= 0; slot--:
        if slots[slot] is empty:
            slots[slot] = task
            break
```
</details>

---

## 📊 Complexity

| Pattern | Time | Space |
|---------|------|-------|
| Jump Game | O(n) | O(1) |
| Interval Scheduling | O(n log n) | O(1) |
| Gas Station | O(n) | O(1) |
| Partition Labels | O(n) | O(1) |

---

## ⚠️ Common Pitfalls

1. **Assuming greedy works**: Always verify! Some problems NEED DP
2. **Wrong greedy choice**: Picking largest first isn't always right
3. **Edge cases**: Single element, already sorted, all overlaps
4. **Counterexamples**: Before implementing, try to find a case where greedy fails

---

## 🚀 When to Use Greedy

✅ **Greedy might work when:**
- Problem has optimal substructure
- Local choices lead to global optimum
- You can PROVE it (exchange argument)
- Classic problems with known greedy solutions

❌ **Don't use Greedy when:**
- Need to explore all possibilities (use backtracking)
- Need optimal substructure but greedy choice fails (use DP)
- Problem is NP-hard (greedy won't give optimal)

---

## 💡 Proving Greedy Works

**Exchange Argument**: Show that any optimal solution can be transformed to use the greedy choice without making it worse.

Example for interval scheduling:
1. Let G be the greedy choice (earliest finishing interval)
2. Let O be an optimal solution
3. If O doesn't use G, replace O's first interval with G
4. This doesn't make O worse (G ends earlier or same time)
5. Therefore, greedy leads to optimal solution

---

## 🧩 Famous Greedy Algorithms

| Algorithm | Problem | Greedy Choice |
|-----------|---------|--------------|
| Dijkstra | Shortest path | Always pick closest unvisited vertex |
| Prim's | MST | Always add cheapest edge to tree |
| Kruskal's | MST | Always add cheapest edge (if no cycle) |
| Huffman | Compression | Combine two least frequent symbols |

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[13-intervals|Intervals]] — Interval scheduling is a classic greedy problem
- [[08-dp-1d|DP 1D]] — Greedy vs DP: local optimal vs global optimal
- [[05-binary-search|Binary Search]] — Binary search on answer uses greedy
- [[18-advanced-graphs|Advanced Graphs]] — Prim's/Kruskal's are greedy MST algorithms
