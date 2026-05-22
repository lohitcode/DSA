---
title: "Pattern: Arrays and Hashing"
aliases: [Pattern: Arrays and Hashing]
type: pattern
tags:
  - pattern
---

# Arrays and Hashing

## Concept
**Arrays** provide contiguous memory storage with O(1) indexed access. **Hashing** uses hash tables (maps/sets) to achieve O(1) average-case lookups, inserts, and deletes by trading space for time.

## Internal Structure

### Array Memory Layout
```
Memory Address:  0x100    0x104    0x108    0x10C
                 ┌────────┬────────┬────────┬────────┐
                 │   10   │   25   │   33   │   47   │
                 └────────┴────────┴────────┴────────┘
Index:              [0]      [1]      [2]      [3]
```
- **Contiguous**: Elements stored sequentially in memory
- **Cache-friendly**: Spatial locality improves performance
- **Fixed size**: Resizing requires O(n) copy operation

### Hash Table Internals
```
Hash Function → Bucket Index → [Linked List of Entries]
     key.hashCode() % 16              ↓
                            ┌─────────────────────┐
                            │ key1 → val1 → nil   │
                            │ key2 → val2 → nil   │
                            │ key3 → val3 → val4  │ (collision chain)
                            └─────────────────────┘
```
- **Hash function**: Maps key to bucket index
- **Collision resolution**: Chaining (linked lists) or Open Addressing
- **Load factor**: `size / capacity` — triggers rehash when exceeded (typically 0.75)
- **Rehashing**: O(n) operation to redistribute all entries

## Common Patterns

### 1. Contains Duplicate
**Pattern**: Use set to track seen elements.
**When**: Need to detect duplicates in O(n) time.

```go
seen := make(map[int]bool)
for _, num := range nums {
    if seen[num] { return true }
    seen[num] = true
}
```

### 2. Valid Anagram
**Pattern**: Count character frequencies.
**When**: Need to compare if two strings have same characters.

```go
// Method 1: Count array (O(1) space for fixed alphabet)
count := [26]int{}
for _, ch := range s { count[ch-'a']++ }
for _, ch := range t { count[ch-'a']-- }

// Method 2: Sorting (O(n log n) time, O(1) space)
sort.Strings([]string{s, t})
return s == t
```

### 3. Two Sum
**Pattern**: Use map for complement lookup in one pass.
**When**: Find pairs that sum to target.

```go
seen := make(map[int]int)  // value → index
for i, num := range nums {
    if j, exists := seen[target-num]; exists {
        return []int{j, i}
    }
    seen[num] = i
}
```

### 4. Group Anagrams
**Pattern**: Use sorted string or count array as key.

```go
groups := make(map[string][]string)
for _, word := range words {
    key := sortString(word)  // Or use count array
    groups[key] = append(groups[key], word)
}
```

## Complexity Analysis

| Operation | Array | Hash Map | Notes |
|-----------|-------|----------|-------|
| Access by index | O(1) | O(1) avg | Array: direct offset; Hash: compute + lookup |
| Search | O(n) | O(1) avg | Hash: O(n) worst case (all collisions) |
| Insert at end | O(1) amortized | O(1) avg | Array: O(n) when resize needed |
| Delete | O(n) | O(1) avg | Array: requires shift |
| Space | O(n) | O(n) | Hash uses ~2-3x more memory |

## Hash Map Performance Factors
1. **Load factor**: Higher = more collisions, slower lookups
2. **Hash quality**: Poor distribution = more collisions
3. **Bucket structure**: Linked lists vs balanced trees (Java 8+ uses trees for >8 entries)
4. **Resize policy**: Triggering rehash too often = performance hit

## Common Pitfalls
1. **Hashing mutable objects**: Don't use slices as map keys in Go (not hashable)
2. **Floating point keys**: Precision issues with `float64` as keys
3. **Modifying while iterating**: Go allows this but behavior undefined
4. **Zero values**: `nil` vs empty map — `nil[v]` returns zero value without error
5. **Order**: Maps are unordered — use `sync.Map` or external ordering if needed

## Go Map Implementation Details
```go
// Set using map (value is struct{} for zero memory)
seen := make(map[int]struct{})
seen[num] = struct{}{}

// Check if exists (comma-ok idiom)
if _, exists := m[key]; exists {
    // key exists
}

// Delete from map
delete(m, key)

// Iterate (order NOT guaranteed)
for k, v := range m {
    // Process k, v
}
```

**Under the hood**: Go maps use:
- **Buckets**: Array of 8-entry buckets
- **Top hash**: First 8 bits of hash stored for quick comparison
- **Overflow**: Extra buckets for collisions (linked list)

## When to Use

### Use Arrays when:
- Need indexed access
- Know size upfront
- Memory is constrained
- Need cache-friendly iteration

### Use Hash Maps when:
- Need fast lookups by key
- Don't know keys in advance
- Need to detect/count/track elements
- Need set operations

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[21-design/21-design-problems|Design Problems]] — LRU Cache uses hash map for O(1) lookups
- [[15-trie|Trie]] — Alternative to hash maps for prefix-based operations
- [[12-heap|Heap]] — Compare with hash map for priority operations
- [[01-linked-list|Linked List]] — Compare memory layout and access patterns
- [[17-bit-manipulation|Bit Manipulation]] — Hash functions use bit operations
