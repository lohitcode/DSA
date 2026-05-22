---
title: "Pattern: Tree DFS"
aliases: [Pattern: Tree DFS]
type: pattern
tags:
  - pattern
---

# Tree DFS (Depth-First Search)

## 🎯 The Core Idea

Imagine exploring a maze. You go **deep** into one path until you hit a dead end, then backtrack and try another path. That's DFS — **go deep before wide**.

```
        1
       / \
      2   3
     / \
    4   5

DFS path: 1 → 2 → 4 → 5 → 3
        (deep → backtrack → deep)
```

**> Quick thought**: Why would you prefer DFS over BFS?

<details>
<summary>Click to reveal...</summary>

DFS uses **less memory** (just the call stack) and is simpler to implement recursively. Use it when you need to explore paths end-to-end.
</details>

---

## 🧠 Tree Traversal Orders (The Big Three)

### Pre-Order: Root → Left → Right
**> "Visit the node FIRST, then children"**

```go
func preorder(node *TreeNode) {
    if node == nil { return }
    visit(node)      // ← First!
    preorder(node.Left)
    preorder(node.Right)
}
```

**When to use**: Creating a copy of tree, serializing tree

---

### In-Order: Left → Root → Right
**> "Visit BETWEEN children"**

```go
func inorder(node *TreeNode) {
    if node == nil { return }
    inorder(node.Left)
    visit(node)      // ← Middle!
    inorder(node.Right)
}
```

**When to use**: BST gives sorted order! Find elements in range.

---

### Post-Order: Left → Right → Root
**> "Visit AFTER children"**

```go
func postorder(node *TreeNode) {
    if node == nil { return }
    postorder(node.Left)
    postorder(node.Right)
    visit(node)      // ← Last!
}
```

**When to use**: Delete subtree, calculate size/height bottom-up

**> 💡 Pro tip**: Post-order is perfect when you need children's results BEFORE processing parent.

---

## 📝 Visual Cheat Sheet

```
Tree:       1
           / \
          2   3

Pre-Order:  1, 2, 3    (Root first)
In-Order:   2, 1, 3    (Root middle)  
Post-Order: 2, 3, 1    (Root last)
```

---

## 🔥 Common Tree Problems (With Patterns)

### Pattern 1: "Bottom-Up" Calculation
**> I need info from children before parent**

**Examples**: Max depth, height, diameter, is balanced

```go
func maxDepth(node *TreeNode) int {
    if node == nil { return 0 }
    
    left := maxDepth(node.Left)   // Get child info first
    right := maxDepth(node.Right)
    
    return 1 + max(left, right)   // Then use it
}
```

**Key insight**: Return value FROM recursive call tells you about subtree.

---

### Pattern 2: "Top-Down" Passing
**> I need to pass info TO children**

**Examples**: Path sum, ancestor info, accumulation

```go
func hasPathSum(node *TreeNode, target int) bool {
    if node == nil { return false }
    
    target -= node.Val  // Consume current node
    
    if node.Left == nil && node.Right == nil {
        return target == 0  // Leaf: did we hit target?
    }
    
    return hasPathSum(node.Left, target) ||
           hasPathSum(node.Right, target)
}
```

**Key insight**: Add a parameter to pass info down the recursion.

---

### Pattern 3: "Global State" Tracking
**> I need to remember something across ALL calls**

**Examples**: Max path sum, diameter, LCA

```go
var maxSum int  // ← Global variable

func maxPathSum(node *TreeNode) int {
    if node == nil { return 0 }
    
    left := max(0, maxPathSum(node.Left))
    right := max(0, maxPathSum(node.Right))
    
    // Update global best
    maxSum = max(maxSum, node.Val + left + right)
    
    return node.Val + max(left, right)  // Path through one side
}
```

**Key insight**: Use global variable or return a struct with multiple values.

---

## 🎮 Practice Exercise

**> Problem**: Find the diameter of a binary tree (longest path between any two nodes).

<details>
<summary>Hint 1: What do I need from each node?</summary>

From each node, I need: the height of its left and right subtrees.
</details>

<details>
<summary>Hint 2: Where can the longest path be?</summary>

Either: entirely in left subtree, OR entirely in right, OR passing through root (leftHeight + rightHeight)
</details>

<details>
<summary>Solution</summary>

```go
var diameter int

func height(node *TreeNode) int {
    if node == nil { return 0 }
    
    left := height(node.Left)
    right := height(node.Right)
    
    // Path through this node
    diameter = max(diameter, left + right)
    
    return 1 + max(left, right)
}
```
</details>

---

## 🔄 Iterative DFS (When Recursion Fails)

**> Problem: Stack overflow on deep trees!**

```go
func preorderIterative(root *TreeNode) {
    if root == nil { return }
    
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]  // Peek
        stack = stack[:len(stack)-1] // Pop
        
        visit(node)
        
        // Push right first, then left (left processed first)
        if node.Right != nil { stack = append(stack, node.Right) }
        if node.Left != nil { stack = append(stack, node.Left) }
    }
}
```

**Key insight**: Stack reverses order — push opposite of desired order.

---

## ⚠️ Common Pitfalls

1. **Forgot nil check**: Always first line — `if node == nil { return }`
2. **Wrong traversal order**: Trace through small tree manually
3. **Not handling leaf**: Leaves are `node.Left == nil && node.Right == nil`
4. **Stack overflow**: Use iterative for very deep trees (like linked list masquerading as tree)

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(n) | Visit every node exactly once |
| Space | O(h) | h = height (call stack), O(n) worst case (skewed tree) |

**Skewed tree**: Like a linked list — each node has only one child

---

## 🚀 When to Use DFS vs BFS

| Use DFS when... | Use BFS when... |
|-----------------|-----------------|
| Need to explore paths fully | Need shortest path |
| Memory is constrained | Need level-by-level processing |
| Problem naturally recursive | Need to find closest neighbor |
| Processing bottom-up | Processing top-down by level |

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[09-tree-bfs|Tree BFS]] — DFS (deep first) vs BFS (level by level)
- [[10-graphs-dfs|Graph DFS]] — Tree DFS is simpler (no cycles)
- [[19-bst/19-binary-search-tree|Binary Search Tree]] — BST operations use DFS
- [[06-stack|Stack]] — DFS uses recursion stack (or explicit stack)
- [[11-backtracking|Backtracking]] — Uses DFS to explore all possibilities
