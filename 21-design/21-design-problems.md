---
title: "Pattern: Design Problems"
aliases: [Pattern: Design Problems]
type: pattern
tags:
  - pattern
---

# Design Problems

## 🎯 The Core Idea

Design problems ask you to **create a data structure** that supports specific operations efficiently. The key is choosing the right **combination** of underlying structures.

**> Quick thought**: Why not just use one data structure for everything?

<details>
<summary>Click to reveal...</summary>

Every structure has trade-offs. Arrays have O(1) access but O(n) insertion. Hash maps have O(1) lookup but no order. Design problems require combining structures to get the best of both worlds.
</details>

---

## 🧠 Design Framework

1. **Understand requirements**: What operations? What frequency?
2. **Analyze trade-offs**: Time vs. space, simplicity vs. performance
3. **Choose primitives**: Arrays, maps, sets, stacks, queues, linked lists
4. **Combine intelligently**: Each structure handles what it's best at

---

## 🔥 Common Design Patterns

### Pattern 1: LRU Cache (Least Recently Used)
**> Evict least recently used item when at capacity**

**Combination**: Hash Map + Doubly Linked List

- **Hash map**: O(1) access to any node
- **Doubly linked list**: O(1) move to front/back

```go
type LRUCache struct {
    capacity int
    cache    map[int]*DListNode
    head, tail *DListNode
}

type DListNode struct {
    key, val int
    prev, next *DListNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        capacity: capacity,
        cache: make(map[int]*DListNode),
    }
    // Dummy head and tail simplify edge cases
    lru.head = &DListNode{}
    lru.tail = &DListNode{}
    lru.head.next = lru.tail
    lru.tail.prev = lru.head
    return lru
}

func (lru *LRUCache) Get(key int) int {
    if node, ok := lru.cache[key]; ok {
        lru.moveToFront(node)
        return node.val
    }
    return -1
}

func (lru *LRUCache) Put(key int, value int) {
    if node, ok := lru.cache[key]; ok {
        node.val = value
        lru.moveToFront(node)
    } else {
        node := &DListNode{key: key, val: value}
        lru.cache[key] = node
        lru.addToFront(node)
        
        if len(lru.cache) > lru.capacity {
            lru.removeLRU()
        }
    }
}

func (lru *LRUCache) moveToFront(node *DListNode) {
    lru.remove(node)
    lru.addToFront(node)
}

func (lru *LRUCache) remove(node *DListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (lru *LRUCache) addToFront(node *DListNode) {
    node.next = lru.head.next
    node.prev = lru.head
    lru.head.next.prev = node
    lru.head.next = node
}

func (lru *LRUCache) removeLRU() {
    lru.tail.prev.key // LRU node
    lru.remove(lru.tail.prev)
    delete(lru.cache, lru.tail.prev.key)
}
```

**Why doubly linked list?** Need to remove from middle (LRU) and move to front — both O(1) with doubly linked.

---

### Pattern 2: Min/Max Stack
**> Stack that can return minimum/maximum in O(1)**

**Idea**: Auxiliary stack tracks min/max at each level.

```go
type MinStack struct {
    stack []int
    mins  []int  // Parallel stack for minimums
}

func Constructor() MinStack {
    return MinStack{}
}

func (s *MinStack) Push(val int) {
    s.stack = append(s.stack, val)
    
    if len(s.mins) == 0 {
        s.mins = append(s.mins, val)
    } else {
        s.mins = append(s.mins, min(val, s.mins[len(s.mins)-1]))
    }
}

func (s *MinStack) Pop() {
    s.stack = s.stack[:len(s.stack)-1]
    s.mins = s.mins[:len(s.mins)-1]
}

func (s *MinStack) Top() int {
    return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
    return s.mins[len(s.mins)-1]
}
```

**Key insight**: `mins[i]` = minimum in `stack[0:i+1]`. When we pop, both stacks stay in sync.

---

### Pattern 3: Queue from Stacks
**> Implement FIFO queue using LIFO stacks**

**Idea**: Two stacks — one for pushing, one for popping.

```go
type MyQueue struct {
    in, out []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (q *MyQueue) Push(x int) {
    q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
    q Peek()
    val := q.out[len(q.out)-1]
    q.out = q.out[:len(q.out)-1]
    return val
}

func (q *MyQueue) Peek() int {
    if len(q.out) == 0 {
        // Transfer all from in to out (reverses order)
        for len(q.in) > 0 {
            val := q.in[len(q.in)-1]
            q.in = q.in[:len(q.in)-1]
            q.out = append(q.out, val)
        }
    }
    return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
    return len(q.in) == 0 && len(q.out) == 0
}
```

**Amortized O(1)**: Each element moved exactly once (from `in` to `out`).

---

### Pattern 4: Trie (Prefix Tree)
**> Efficient prefix-based string operations**

See the **Trie** pattern for full implementation.

**Operations**:
- Insert word: O(m) where m = word length
- Search word: O(m)
- StartsWith prefix: O(m)

---

### Pattern 5: Insert Delete GetRandom O(1)
**> Data structure with insert, delete, getRandom in O(1)**

**Combination**: Array + Hash Map

```go
type RandomizedSet struct {
    nums   []int
    indices map[int]int  // value → index in nums
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        indices: make(map[int]int),
    }
}

func (rs *RandomizedSet) Insert(val int) bool {
    if _, exists := rs.indices[val]; exists {
        return false
    }
    rs.indices[val] = len(rs.nums)
    rs.nums = append(rs.nums, val)
    return true
}

func (rs *RandomizedSet) Remove(val int) bool {
    idx, exists := rs.indices[val]
    if !exists {
        return false
    }
    
    // Swap with last element
    last := rs.nums[len(rs.nums)-1]
    rs.nums[idx] = last
    rs.indices[last] = idx
    
    // Remove last
    rs.nums = rs.nums[:len(rs.nums)-1]
    delete(rs.indices, val)
    
    return true
}

func (rs *RandomizedSet) GetRandom() int {
    return rs.nums[rand.Intn(len(rs.nums))]
}
```

**Key insight**: Array gives O(1) random access. Hash map gives O(1) find. Swap-to-delete makes removal O(1).

---

## 🎮 Practice Exercise

**> Design a data structure that supports insert, remove, and getRandom with duplicates allowed.**

<details>
<summary>Hint: How to handle duplicates?</summary>

Instead of mapping value → index, map value → SET of indices.
</details>

<details>
<summary>Approach</summary>

```go
type RandomizedCollection struct {
    nums   []int
    indices map[int]map[int]struct{}  // value → set of indices
}

func (rc *RandomizedCollection) Insert(val int) bool {
    exists := len(rc.indices[val]) > 0
    rc.indices[val][len(rc.nums)] = struct{}{}
    rc.nums = append(rc.nums, val)
    return !exists
}

func (rc *RandomizedCollection) Remove(val int) bool {
    if len(rc.indices[val]) == 0 {
        return false
    }
    
    // Get any index of val
    idx := anyKey(rc.indices[val])
    last := rc.nums[len(rc.nums)-1]
    
    // Move last to idx
    rc.nums[idx] = last
    delete(rc.indices[val], idx)
    rc.indices[last][idx] = struct{}{}
    delete(rc.indices[last], len(rc.nums)-1)
    
    rc.nums = rc.nums[:len(rc.nums)-1]
    return true
}
```
</details>

---

## 📊 Design Pattern Cheat Sheet

| Requirement | Structure Combination |
|-------------|---------------------|
| Fast lookup + order | Hash map + Linked list |
| LRU eviction | Hash map + Doubly linked list |
| Min/Max operations | Main stack + Aux stack |
| Random access + Delete | Array + Hash map |
| Prefix operations | Trie |
| Range queries | Segment tree / Binary indexed tree |

---

## ⚠️ Common Pitfalls

1. **Thread safety**: Most designs aren't thread-safe (add locks if needed)
2. **Memory leaks**: Properly clean up when removing
3. **Edge cases**: Empty structure, single element, duplicates
4. **Amortized vs. worst case**: Be clear about time complexity
5. **Space trade-offs**: Optimizing time often costs more space

---

## 🚀 Design Heuristics

1. **Start simple**: Naive implementation first, optimize later
2. **Identify bottlenecks**: Which operation is called most?
3. **Hot path optimization**: Make frequent operations fastest
4. **Separation of concerns**: Each structure has one responsibility
5. **Test edge cases**: Empty, single element, max capacity

---

## 💡 System Design vs. Data Structure Design

| System Design | Data Structure Design |
|---------------|----------------------|
| Scalability, distributed | Time/space complexity |
| Databases, caching | Arrays, maps, trees |
| Load balancers, queues | Stacks, heaps, tries |
| Real-world constraints | Algorithm requirements |

Both follow the same principle: **understand requirements, choose right tools**.

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[01-linked-list|Linked List]] — LRU Cache uses doubly linked list
- [[03-arrays-hashing|Arrays & Hashing]] — LRU Cache uses hash map
- [[06-stack|Stack]] — Min Stack, Queue from Stacks
- [[12-heap|Heap]] — Priority queue design
