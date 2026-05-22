---
title: "Pattern: Tree BFS"
aliases: [Pattern: Tree BFS]
type: pattern
tags:
  - pattern
---

# Tree BFS (Breadth-First Search)

## 🎯 The Core Idea

Imagine you're exploring a tree **level by level** — like reading a book from top to bottom, left to right. That's BFS: explore all nodes at depth 1, then depth 2, then depth 3, and so on.

```
        1         Level 0: [1]
       / \
      2   3       Level 1: [2, 3]
     / \
    4   5         Level 2: [4, 5]

BFS order: 1 → 2 → 3 → 4 → 5
```

**> Quick thought**: When would you prefer BFS over DFS?

<details>
<summary>Click to reveal...</summary>

Use BFS when you need the **shortest path** or need to process nodes **level by level**. DFS goes deep first, so it might find a longer path before finding the shortest one.
</details>

---

## 🧠 How BFS Works (The Queue)

BFS uses a **queue** (FIFO — First In, First Out):

```go
queue := []*TreeNode{root}

for len(queue) > 0 {
    node := queue[0]      // Peek front
    queue = queue[1:]     // Dequeue
    
    // Process node...
    
    if node.Left != nil {
        queue = append(queue, node.Left)   // Enqueue left
    }
    if node.Right != nil {
        queue = append(queue, node.Right)  // Enqueue right
    }
}
```

**Key insight**: Queue ensures we process nodes in order they were discovered — which means level by level!

---

## 🔥 Common BFS Patterns

### Pattern 1: Level Order Traversal
**> Return nodes grouped by level**

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    
    result := [][]int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []int{}
        
        // Process exactly one level
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            currentLevel = append(currentLevel, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, currentLevel)
    }
    
    return result
}
```

**Key insight**: Capture queue size BEFORE processing level — this tells you how many nodes are at current level.

---

### Pattern 2: Zigzag Level Order
**> Alternate direction each level**

```go
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    
    result := [][]int{}
    queue := []*TreeNode{root}
    leftToRight := true
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            // Fill based on direction
            if leftToRight {
                currentLevel[i] = node.Val
            } else {
                currentLevel[levelSize-1-i] = node.Val
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, currentLevel)
        leftToRight = !leftToRight  // Flip direction
    }
    
    return result
}
```

---

### Pattern 3: Level Averages
**> Compute average of each level**

```go
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil { return []float64{} }
    
    result := []float64{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        sum := 0
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            sum += node.Val
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, float64(sum)/float64(levelSize))
    }
    
    return result
}
```

---

### Pattern 4: Right Side View
**> What would you see from the right?**

```go
func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }
    
    result := []int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            // Last node in level = visible from right
            if i == levelSize-1 {
                result = append(result, node.Val)
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return result
}
```

**Key insight**: The rightmost node at each level is what you see.

---

### Pattern 5: Shortest Path in Binary Tree
**> Minimum steps from root to target value**

```go
func findTarget(root *TreeNode, target int) int {
    if root == nil { return -1 }
    if root.Val == target { return 0 }
    
    queue := []*TreeNode{root}
    steps := 0
    
    for len(queue) > 0 {
        steps++
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left != nil {
                if node.Left.Val == target { return steps }
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                if node.Right.Val == target { return steps }
                queue = append(queue, node.Right)
            }
        }
    }
    
    return -1  // Not found
}
```

**Key insight**: BFS guarantees shortest path in unweighted graphs/trees!

---

## 🎮 Practice Exercise

**> Problem**: Find the largest value in each level of a binary tree.

<details>
<summary>Think about what to track...</summary>

Track the maximum value while processing each level.
</details>

<details>
<summary>Solution</summary>

```go
func largestValues(root *TreeNode) []int {
    if root == nil { return []int{} }
    
    result := []int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        maxVal := math.MinInt
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            maxVal = max(maxVal, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, maxVal)
    }
    
    return result
}
```
</details>

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(n) | Visit every node exactly once |
| Space | O(w) | w = max width of tree (max queue size) |

**Worst case**: Complete binary tree has width n/2 at last level

---

## ⚠️ Common Pitfalls

1. **Wrong level tracking**: Capture queue size BEFORE inner loop
2. **Enqueue order**: Usually left first, then right (unless problem says otherwise)
3. **Nil check**: Always check if root is nil first
4. **Queue indexing**: `queue[0]` to peek, `queue[1:]` to dequeue (or use slice with index)
5. **Forgotten nodes**: Make sure to enqueue BOTH children when they exist

---

## 🚀 BFS vs DFS for Trees

| Use BFS when... | Use DFS when... |
|-----------------|-----------------|
| Need shortest path | Need to explore paths deeply |
| Need level-by-level processing | Need simple recursive solution |
| Processing top-down | Processing bottom-up |
| Queue is fine | Want less memory (no queue) |

---

## 💡 Optimization Tip

Instead of slicing the queue (O(n) operation), use an index:

```go
// Less efficient:
queue = queue[1:]

// More efficient:
idx := 0
for idx < len(queue) {
    node := queue[idx]
    idx++
    // ... enqueue children
}
```

This avoids reallocation — queue grows but we just move our read pointer.

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[07-tree-dfs|Tree DFS]] — BFS (level by level) vs DFS (deep first)
- [[18-advanced-graphs|Advanced Graphs]] — Dijkstra uses BFS-like approach
- [[12-heap|Heap]] — BFS uses queue, heap is a priority queue
- [[06-stack|Stack]] — BFS queue (FIFO) vs Stack (LIFO)
