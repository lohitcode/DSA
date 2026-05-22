---
title: "Pattern: Backtracking"
aliases: [Pattern: Backtracking]
type: pattern
tags:
  - pattern
---

# Backtracking

## 🎯 The Core Idea

Imagine you're in a maze. You go down one path, hit a dead end, **backtrack** to the last junction, and try another path. Keep doing this until you've explored everything.

That's backtracking: **systematic trial and error** with undo.

```
        Start
       /    \
      A      B     Try path A...
     / \          Hit dead end, backtrack!
    C   D         Try path D... Found exit!
```

**> Quick thought**: How is backtracking different from plain recursion?

<details>
<summary>Click to reveal...</summary>

Backtracking is recursion with **undo**. You make a choice, recurse, then **undo the choice** before trying the next option. This explores ALL possibilities without using extra memory for copies.
</details>

---

## 🧠 The Backtracking Template

```go
func backtrack(choices, currentState) {
    // Base case: Found a valid solution
    if isComplete(currentState) {
        result = append(result, copy(currentState))
        return
    }
    
    // Try each choice
    for eachChoice in availableChoices {
        if isValid(eachChoice, currentState) {
            make(eachChoice)              // Choose
            backtrack(..., currentState)   // Explore
            undo(eachChoice)              // UNDO ← Key!
        }
    }
}
```

**The magic**: The `undo` step lets us reuse memory instead of creating new arrays each time.

---

## 🔥 Common Backtracking Patterns

### Pattern 1: Subsets (Power Set)
**> Generate ALL possible combinations**

```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    backtrack(nums, 0, []int{}, &result)
    return result
}

func backtrack(nums []int, start int, path []int, result *[][]int) {
    // Add copy of current path
    temp := make([]int, len(path))
    copy(temp, path)
    *result = append(*result, temp)
    
    // Try adding each remaining element
    for i := start; i < len(nums); i++ {
        path = append(path, nums[i])      // Choose
        backtrack(nums, i+1, path, result) // Explore
        path = path[:len(path)-1]          // Undo
    }
}
```

**Example**: `[1, 2]` → `[], [1], [2], [1, 2]`

**Key insight**: At each element, you have 2 choices: include it or skip it.

---

### Pattern 2: Permutations
**> Generate ALL arrangements**

```go
func permute(nums []int) [][]int {
    result := [][]int{}
    backtrack(nums, []int{}, &result)
    return result
}

func backtrack(remaining, path []int, result *[][]int) {
    // Base case: all numbers used
    if len(remaining) == 0 {
        temp := make([]int, len(path))
        copy(temp, path)
        *result = append(*result, temp)
        return
    }
    
    // Try each remaining number
    for i := 0; i < len(remaining); i++ {
        // Choose: add to path, remove from remaining
        path = append(path, remaining[i])
        newRemaining := append(append([]int{}, remaining[:i]...), remaining[i+1:]...)
        
        backtrack(newRemaining, path, result)
        
        // Undo
        path = path[:len(path)-1]
    }
}
```

**Example**: `[1, 2, 3]` → `[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]`

**Optimization**: Use swap-in-place to avoid copying `remaining` array.

---

### Pattern 3: Combinations (n Choose k)
**> Generate k-sized groups**

```go
func combine(n int, k int) [][]int {
    result := [][]int{}
    backtrack(1, n, k, []int{}, &result)
    return result
}

func backtrack(start, n, k int, path []int, result *[][]int) {
    // Base case: found k elements
    if len(path) == k {
        temp := make([]int, len(path))
        copy(temp, path)
        *result = append(*result, temp)
        return
    }
    
    // Pruning: Not enough elements left?
    // remaining elements needed: k - len(path)
    // available elements: n - i + 1
    for i := start; i <= n - (k - len(path)) + 1; i++ {
        path = append(path, i)
        backtrack(i+1, n, k, path, result)
        path = path[:len(path)-1]
    }
}
```

**Example**: `n=4, k=2` → `[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]`

---

### Pattern 4: With Duplicates (Skip Repeats)
**> Avoid duplicate combinations**

```go
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)  // Sort first!
    result := [][]int{}
    backtrack(candidates, target, 0, []int{}, &result)
    return result
}

func backtrack(candidates []int, target, start int, path []int, result *[][]int) {
    if target == 0 {
        temp := make([]int, len(path))
        copy(temp, path)
        *result = append(*result, temp)
        return
    }
    
    for i := start; i < len(candidates); i++ {
        // Skip duplicates at same level
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }
        
        if candidates[i] > target {
            break  // Sorted, so no more valid
        }
        
        path = append(path, candidates[i])
        backtrack(candidates, target-candidates[i], i+1, path, result)
        path = path[:len(path)-1]
    }
}
```

**Key insight**: Sort first, then skip `candidates[i] == candidates[i-1]` at same recursion level.

---

## 🎮 Practice Exercise

**> Problem**: Generate all valid parentheses combinations with n pairs.

**Example**: n=2 → `(())`, `()()`

<details>
<summary>Hint: When can you add '(' vs ')'?</summary>

You can always add '(' (if not exceeded n). You can add ')' only if more '(' than ')' in current string.
</details>

<details>
<summary>Solution</summary>

```go
func generateParenthesis(n int) []string {
    result := []string{}
    backtrack(n, 0, 0, "", &result)
    return result
}

func backtrack(n, open, close int, path string, result *[]string) {
    // Base case: complete
    if len(path) == 2*n {
        *result = append(*result, path)
        return
    }
    
    // Add '(' if we haven't used all
    if open < n {
        backtrack(n, open+1, close, path+"(", result)
    }
    
    // Add ')' if we have more '(' than ')'
    if close < open {
        backtrack(n, open, close+1, path+")", result)
    }
}
```
</details>

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(k × 2ⁿ) or O(k × n!) | Generate all combinations/permutations |
| Space | O(n) | Recursion depth + current path |

n = input size, k = average length of solution

**Note**: This is EXPENSIVE! Only use when you must explore all possibilities.

---

## ⚠️ Common Pitfalls

1. **Forgot to copy**: `result = append(result, path)` copies reference! Use `copy()`
2. **Wrong undo**: Must undo exactly what you did
3. **Missing pruning**: Add `if` conditions to skip invalid branches early
4. **Duplicate handling**: Sort + skip duplicates at same level
5. **Off-by-one**: Careful with start indices in loops

---

## 🚀 When to Use Backtracking

✅ **Use Backtracking when:**
- Need to generate ALL possibilities (combinations, permutations)
- Need ALL valid solutions
- Problem can be broken into choices
- Can prune invalid branches early

❌ **Don't use Backtracking when:**
- Just need ONE solution (might be faster algorithm)
- Input size is large (exponential time!)
- Problem has greedy/optimal substructure

---

## 💡 Optimization: Pruning

**Pruning** = cutting branches early that can't possibly lead to valid solutions.

```go
// BEFORE pruning: explore everything
for i := start; i < len(nums); i++ {
    backtrack(...)
}

// AFTER pruning: skip impossible branches
for i := start; i <= n - needed + 1; i++ {
    if invalidCondition(i) { continue }  // Skip!
    backtrack(...)
}
```

**Example**: For combinations, if you need 3 more elements but only 2 left, stop early.

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[07-tree-dfs|Tree DFS]] — Backtracking uses DFS to explore all possibilities
- [[10-graphs-dfs|Graph DFS]] — Backtracking on graphs explores all paths
- [[08-dp-1d|DP 1D]] — DP is optimized backtracking with memoization
- [[06-stack|Stack]] — Backtracking uses stack for state management
