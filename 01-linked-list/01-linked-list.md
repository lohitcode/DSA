---
title: "Pattern: Linked List"
aliases: [Pattern: Linked List]
type: pattern
tags:
  - pattern
---

# Linked List

## Concept
A linked list is a linear data structure where elements are stored in nodes, and each node points to the next node via a reference/pointer. Unlike arrays, linked lists allow efficient insertion and deletion at any position.

## Internal Structure
```
┌─────┐    ┌─────┐    ┌─────┐    ┌─────┐
│  1  │───▶│  2  │───▶│  3  │───▶│ nil │
└─────┘    └─────┘    └─────┘    └─────┘
  Val        Val        Val       (end)
  Next       Next       Next
```

**Memory Layout**: Nodes are scattered in memory (non-contiguous), unlike arrays which require contiguous allocation. Each node typically contains:
- `Val`: The actual data
- `Next`: Pointer/reference to the next node (8 bytes on 64-bit systems)

## Go Node Definition
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

## Common Patterns

### 1. Dummy Head Technique
**Why**: Eliminates edge case handling for head modifications (insertion at start, deletion of head).

```go
dummy := &ListNode{0, head}
current := dummy
// ... operate on current.Next
return dummy.Next  // New head (may be different)
```

### 2. Fast/Slow Pointers
**Why**: Detect cycles, find middle element in one pass.
- **Fast**: Moves 2 steps at a time
- **Slow**: Moves 1 step at a time
- **Meeting point**: If cycle exists, fast will catch slow

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast { /* cycle found */ }
}
```

### 3. Reversal (Iterative)
**Key insight**: Maintain `prev`, `curr`, `next` pointers.

```go
var prev *ListNode
curr := head
for curr != nil {
    next := curr.Next  // Save next
    curr.Next = prev   // Reverse link
    prev = curr        // Advance prev
    curr = next        // Advance curr
}
return prev  // New head
}
```

## Complexity Analysis
| Operation | Linked List | Array | Notes |
|-----------|-------------|-------|-------|
| Access by index | O(n) | O(1) | Must traverse |
| Insert at head | O(1) | O(n) | Array needs shift |
| Delete at head | O(1) | O(n) | Array needs shift |
| Insert at end* | O(n) | O(1) amortized | Need tail pointer for O(1) |
| Delete given node | O(1)* | O(n) | Only if you have ref to node |

*If you have reference to the node and it's not the tail, you can copy next node's value and skip it.

## Common Pitfalls
1. **Off-by-one errors**: Loop conditions (`curr != nil` vs `curr.Next != nil`)
2. **Lost references**: Always save `curr.Next` before modifying `curr.Next`
3. **Cycle in infinite loop**: Always handle termination condition
4. **Memory leaks**: In languages with manual memory management, free deleted nodes
5. **Null pointer dereference**: Check `curr != nil` before accessing `curr.Val`

## When to Use
- Frequent insertions/deletions at head
- Unknown number of elements (dynamic size)
- Don't need random access by index
- Memory efficiency when elements are large

## Traversal
```go
for head != nil {
    // Process head.Val
    head = head.Next
}
```

## Related Topics
- [[06-stack|Stack]] — Linked lists are used to implement stacks
- [[21-design/21-design-problems|Design Problems]] — LRU Cache uses doubly linked list + hash map
- [[03-arrays-hashing|Arrays & Hashing]] — Compare memory layout and access patterns
- [[02-two-pointers|Two Pointers]] — Fast/slow pointers work on linked lists

---
**[[../index.md|← Back to Topics]]**
