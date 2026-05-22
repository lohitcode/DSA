---
title: "Pattern: Heap"
aliases: [Pattern: Heap]
type: pattern
tags:
  - pattern
---

# Heap (Priority Queue)

## 🎯 The Core Idea

Imagine you have a pile of tasks, but some are **more important** than others. A heap lets you quickly grab the **most important** item — like a VIP queue.

**> Quick thought**: Why not just sort an array and grab the first element?

<details>
<summary>Click to reveal...</summary>

Sorting takes O(n log n) ONCE. Heap lets you keep adding/removing items efficiently — O(log n) per operation. Use it when data is **dynamic**.
</details>

---

## 🧠 How a Heap Works (Visual)

A heap is a **complete binary tree** where parents compare to children in a specific way:

### Min-Heap: Parent ≤ Children
```
        1           Level 0: Root (minimum)
       / \
      3   2         Level 1: Children ≥ parent
     / \
    5   4           Level 2: Leaves

Array: [1, 3, 2, 5, 4]
Index:  0  1  2  3  4
```

### Max-Heap: Parent ≥ Children
```
        9           Level 0: Root (maximum)
       / \
      7   8         Level 1: Children ≤ parent
     /
    5               Level 2: Leaves

Array: [9, 7, 8, 5]
```

**Key property**: Parent-child relationship in array:
- Parent of `i` = `(i-1) / 2`
- Left child of `i` = `2i + 1`
- Right child of `i` = `2i + 2`

---

## 📝 The Heap Operations

### Push (Insert): O(log n)
```go
// Add to end, then "bubble up"
heap.Push(&h, value)
```

**How it works**:
1. Add element at the end
2. Compare with parent
3. If violates heap property, swap with parent
4. Repeat until property satisfied

```
Push 2 into [1, 3]:
  1. [1, 3, 2]     Add to end
  2. 2 > 1? Swap!  [2, 3, 1]
  3. Done ✓
```

### Pop (Extract Min/Max): O(log n)
```go
value := heap.Pop(&h)
```

**How it works**:
1. Save root (min/max value)
2. Move last element to root
3. "Bubble down" by swapping with smaller/larger child
4. Repeat until property satisfied

```
Pop from [1, 3, 2]:
  1. Save 1 (answer)
  2. Move 2 to root: [2, 3]
  3. 2 > 3? Swap!   [3, 2]
  4. Return 1 ✓
```

### Peek: O(1)
```go
min := h[0]  // Just look at root
```

---

## 🔥 Common Heap Patterns

### Pattern 1: Top K Elements
**> Find the K largest/smallest elements**

```go
// For K LARGEST: use min-heap of size K
func topKFrequent(nums []int, k int) []int {
    // Count frequencies
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }
    
    // Min-heap of size k
    h := &MinHeap{}
    heap.Init(h)
    
    for num, freq := range count {
        heap.Push(h, Item{num, freq})
        if h.Len() > k {
            heap.Pop(h)  // Remove least frequent
        }
    }
    
    // Extract results
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = heap.Pop(h).(Item).num
    }
    return result
}
```

**Key insight**: Heap size stays at K, so we only keep the K best elements.

---

### Pattern 2: Merge K Sorted Lists
**> Merge multiple sorted arrays efficiently**

```go
func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
    heap.Init(h)
    
    // Push first node of each list
    for _, head := range lists {
        if head != nil {
            heap.Push(h, head)
        }
    }
    
    dummy := &ListNode{}
    curr := dummy
    
    for h.Len() > 0 {
        node := heap.Pop(h).(*ListNode)
        curr.Next = node
        curr = curr.Next
        
        if node.Next != nil {
            heap.Push(h, node.Next)  // Push next from same list
        }
    }
    
    return dummy.Next
}
```

**Key insight**: Always grab the smallest available element across all lists.

---

### Pattern 3: Stream Processing
**> Find median of a data stream**

```go
type MedianFinder struct {
    maxHeap *MaxHeap  // Left half (smaller numbers)
    minHeap *MinHeap  // Right half (larger numbers)
}

func (mf *MedianFinder) AddNum(num int) {
    // Add to appropriate heap
    if mf.maxHeap.Len() == 0 || num <= mf.maxHeap.Peek() {
        heap.Push(mf.maxHeap, num)
    } else {
        heap.Push(mf.minHeap, num)
    }
    
    // Balance: maxHeap can have at most 1 more than minHeap
    if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
        heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
    } else if mf.minHeap.Len() > mf.maxHeap.Len() {
        heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    if mf.maxHeap.Len() > mf.minHeap.Len() {
        return float64(mf.maxHeap.Peek())
    }
    return (float64(mf.maxHeap.Peek()) + float64(mf.minHeap.Peek())) / 2
}
```

**Key insight**: Two heaps divide data — maxHeap for lower half, minHeap for upper half.

---

## 🎮 Practice Exercise

**> Problem**: Given an unsorted array, find the Kth largest element.

<details>
<summary>Think about what heap type to use...</summary>

Use a MIN-heap of size K. The root will be the Kth largest.
</details>

<details>
<summary>Solution</summary>

```go
func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)  // Remove smallest
        }
    }
    
    return heap.Pop(h).(int)
}
```
</details>

---

## 📊 Go's container/heap

```go
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }  // ← Determines min/max
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// Usage:
h := &MinHeap{2, 1, 3}
heap.Init(h)
heap.Push(h, 4)
min := heap.Pop(h)
```

**For Max-Heap**: Just flip the `Less` function: `return h[i] > h[j]`

---

## 📊 Complexity

| Operation | Heap | Sorted Array |
|-----------|------|--------------|
| Find min/max | O(1) | O(1) |
| Insert | O(log n) | O(n) |
| Delete min/max | O(log n) | O(n) |
| Build from array | O(n) | O(n log n) |

---

## ⚠️ Common Pitfalls

1. **Wrong heap type**: Min-heap gives smallest, Max-heap gives largest
2. **Forgot heap.Init**: Must initialize after creating
3. **Type assertions**: `Pop()` returns `any` — need type assertion
4. **Empty heap**: Always check `Len() > 0` before `Pop()`
5. **Modifying during iteration**: Don't modify heap while iterating

---

## 🚀 When to Use a Heap

✅ **Use Heap when:**
- Need quick access to min/max element
- Data is constantly being added/removed
- Need "top K" elements
- Streaming data
- Need priority-based processing

❌ **Don't use Heap when:**
- Just need to find min/max once (sort is simpler)
- Need random access (array is better)
- Data is static (sort once is faster)

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[03-arrays-hashing|Arrays & Hashing]] — Heap vs hash map for priority operations
- [[09-tree-bfs|Tree BFS]] — Heap is a priority queue (BFS uses regular queue)
- [[18-advanced-graphs|Advanced Graphs]] — Dijkstra uses min-heap
- [[21-design/21-design-problems|Design Problems]] — Priority queue design patterns
