---
title: "Pattern: Binary Search Tree"
aliases: [Pattern: Binary Search Tree]
type: pattern
tags:
  - pattern
---

# Binary Search Tree (BST)

## 🎯 The Core Idea

A BST is a binary tree with a **special ordering property**: everything in the left subtree is smaller, everything in the right subtree is larger.

```
        8
       / \
      3   10
     / \    \
    1   6    14
       / \   /
      4   7 13

Left < Node < Right
```

**> Quick thought**: Why is this property so powerful?

<details>
<summary>Click to reveal...</summary>

It enables O(log n) search! At each node, you can eliminate half the remaining tree — just like binary search on arrays.
</details>

---

## 🧠 The BST Property

For EVERY node in the tree:
- All values in LEFT subtree < node's value
- All values in RIGHT subtree > node's value
- Both subtrees are ALSO BSTs

**Critical**: The property must hold for the ENTIRE subtree, not just immediate children!

```
    5           This is NOT a BST!
   / \
  1   6
     /
    3        ← 3 is in right subtree but < 5!
```

---

## 🔥 Common BST Patterns

### Pattern 1: Validate BST
**> Is this tree actually a BST?**

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
    if node == nil { return true }
    
    // Check current node against bounds
    if int64(node.Val) <= min || int64(node.Val) >= max {
        return false
    }
    
    // Recurse with updated bounds
    return validate(node.Left, min, int64(node.Val)) &&
           validate(node.Right, int64(node.Val), max)
}
```

**Key insight**: Pass down **allowed ranges**. Left children get tighter upper bounds, right children get tighter lower bounds.

---

### Pattern 2: Kth Smallest Element
**> Find the kth smallest element in BST**

```go
func kthSmallest(root *TreeNode, k int) int {
    var result int
    var inorder func(node *TreeNode)
    
    inorder = func(node *TreeNode) {
        if node == nil || k <= 0 { return }
        
        inorder(node.Left)    // Visit left first
        k--                    // Count current node
        if k == 0 {
            result = node.Val
            return
        }
        inorder(node.Right)   // Then visit right
    }
    
    inorder(root)
    return result
}
```

**Key insight**: In-order traversal of BST visits nodes in **sorted order**!

---

### Pattern 3: Lowest Common Ancestor (LCA)
**> Find deepest node that is ancestor to BOTH nodes**

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if p.Val < root.Val && q.Val < root.Val {
            // Both in left subtree
            root = root.Left
        } else if p.Val > root.Val && q.Val > root.Val {
            // Both in right subtree
            root = root.Right
        } else {
            // Split point: this is LCA
            return root
        }
    }
    return nil
}
```

**Key insight**: Use BST property — no need to search both subtrees!

```
      20
     /  \
   10    30
  /  \
 5   15

LCA(5, 15) = 10
LCA(5, 30) = 20
```

---

### Pattern 4: Convert Sorted Array to BST
**> Build balanced BST from sorted array**

```go
func sortedArrayToBST(nums []int) *TreeNode {
    return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
    if left > right { return nil }
    
    mid := left + (right-left)/2
    root := &TreeNode{Val: nums[mid]}
    
    root.Left = build(nums, left, mid-1)
    root.Right = build(nums, mid+1, right)
    
    return root
}
```

**Key insight**: Middle element = root (ensures balance). Recursively build left and right.

---

### Pattern 5: Floor and Ceiling
**> Find largest value ≤ target (floor) or smallest ≥ target (ceiling)**

```go
func floor(root *TreeNode, target int) int {
    result := -1
    for root != nil {
        if root.Val == target {
            return target
        } else if root.Val > target {
            root = root.Left
        } else {
            result = root.Val  // Potential floor
            root = root.Right  // Try to find closer
        }
    }
    return result
}
```

---

## 🎮 Practice Exercise

**> Problem**: Find the closest value to a target in a BST.

<details>
<summary>Think about how BST property helps...</summary>

At each node, go toward the target. Track the closest value seen so far.
</details>

<details>
<summary>Solution</summary>

```go
func closestValue(root *TreeNode, target float64) int {
    closest := root.Val
    
    for root != nil {
        // Update closest if current is closer
        if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
            closest = root.Val
        }
        
        // Decide direction
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    
    return closest
}
```
</details>

---

## 📊 Complexity

| Operation | Average | Worst | Notes |
|-----------|---------|-------|-------|
| Search | O(log n) | O(n) | Skewed tree = linked list |
| Insert | O(log n) | O(n) | Find position, insert |
| Delete | O(log n) | O(n) | Hardest BST operation! |
| Min/Max | O(log n) | O(n) | Go all the way left/right |

**Balance matters**: BSTs degrade to O(n) when unbalanced. Use AVL or Red-Black trees for guaranteed balance.

---

## ⚠️ Common Pitfalls

1. **Checking only immediate children**: Must verify ENTIRE subtree obeys BST property
2. **Equal values**: Decide if equal goes left or right (be consistent!)
3. **Integer bounds**: Use int64 for min/max to avoid overflow
4. **Nil pointer**: Always check `node == nil` before accessing `.Val`
5. **Deleting nodes**: Three cases — no children, one child, two children

---

## 🚀 BST Operations

### Search
```go
func search(root *TreeNode, target int) bool {
    for root != nil {
        if root.Val == target {
            return true
        } else if target < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return false
}
```

### Insert
```go
func insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    
    if val < root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }
    
    return root
}
```

### Delete (Tricky!)
```go
func delete(root *TreeNode, key int) *TreeNode {
    if root == nil { return nil }
    
    if key < root.Val {
        root.Left = delete(root.Left, key)
    } else if key > root.Val {
        root.Right = delete(root.Right, key)
    } else {
        // Found node to delete
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }
        
        // Two children: find inorder successor
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = delete(root.Right, minNode.Val)
    }
    
    return root
}
```

---

## 💡 BST vs. Other Structures

| Structure | Search | Insert | Ordered | Balanced |
|-----------|--------|--------|---------|----------|
| BST | O(log n) avg | O(log n) avg | ✓ | Maybe |
| AVL Tree | O(log n) | O(log n) | ✓ | Always |
| Hash Table | O(1) | O(1) | ✗ | N/A |
| Array | O(log n) | O(n) | ✓ | N/A |

---

## 🌳 Tree Traversals Reminder

```go
// In-order: Left → Root → Right (gives sorted order for BST!)
func inorder(root *TreeNode) {
    if root == nil { return }
    inorder(root.Left)
    visit(root)
    inorder(root.Right)
}

// Pre-order: Root → Left → Right
func preorder(root *TreeNode) {
    if root == nil { return }
    visit(root)
    preorder(root.Left)
    preorder(root.Right)
}

// Post-order: Left → Right → Root
func postorder(root *TreeNode) {
    if root == nil { return }
    postorder(root.Left)
    postorder(root.Right)
    visit(root)
}
```

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[05-binary-search|Binary Search]] — BST property enables binary search
- [[07-tree-dfs|Tree DFS]] — BST operations use DFS traversal
- [[21-design/21-design-problems|Design Problems]] — BST is a common data structure design
