---
title: "Pattern: Graph DFS"
aliases: [Pattern: Graph DFS]
type: pattern
tags:
  - pattern
---

# Graph DFS (Depth-First Search)

## 🎯 The Core Idea

Graph DFS is like Tree DFS, but with a **twist**: graphs can have **cycles**. Without tracking where you've been, you'll run in circles forever.

**> Quick thought**: What happens if you don't track visited nodes?

<details>
<summary>Click to reveal...</summary>

Infinite loop! You'll keep visiting the same nodes over and over.
</details>

---

## 🧠 The Secret Sauce: Visited Set

```go
visited := make(map[int]bool)

func dfs(node int) {
    if visited[node] { return }  // ← Already been here!
    visited[node] = true          // ← Mark as visited
    
    // Process node...
    
    for _, neighbor := range graph[node] {
        dfs(neighbor)  // Explore neighbors
    }
}
```

**Why it works**: Once visited, we skip — ensures each node processed once.

---

## 📝 Graph Representations (Pick Your Poison)

### Adjacency List (Most Common)
```go
graph := map[int][]int{
    0: {1, 2},
    1: {0, 3},
    2: {0},
    3: {1},
}
```
**Space**: O(V + E) — efficient for sparse graphs

### Adjacency Matrix
```go
matrix := [][]int{
    {0, 1, 1, 0},  // Node 0 connected to 1, 2
    {1, 0, 0, 1},  // Node 1 connected to 0, 3
    // ...
}
```
**Space**: O(V²) — use for dense graphs or quick edge lookup

---

## 🔥 Common Graph Problems

### Pattern 1: Connected Components
**> How many separate "islands" exist?**

```go
func countComponents(n int, edges [][]int) int {
    // Build graph
    graph := buildGraph(n, edges)
    visited := make(map[int]bool)
    count := 0
    
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, graph, visited)
            count++  // New component found!
        }
    }
    return count
}
```

**Key insight**: Each unvisited node = new component. DFS marks all nodes in that component.

---

### Pattern 2: Island Counting (Grid)
**> Count connected regions in 2D grid**

```go
func numIslands(grid [][]byte) int {
    count := 0
    
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {  // Unvisited land
                dfs(grid, i, j)
                count++
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i, j int) {
    // Boundary / water check
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return
    }
    if grid[i][j] != '1' { return }
    
    grid[i][j] = '0'  // Mark visited (turn to water)
    
    // Explore 4 directions
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j-1)
}
```

**Key insight**: Modify grid in-place to mark visited (saves memory).

---

### Pattern 3: Path Finding
**> Is there a path from A to B?**

```go
func hasPath(graph map[int][]int, start, end int) bool {
    visited := make(map[int]bool)
    return dfsPath(graph, start, end, visited)
}

func dfsPath(graph map[int][]int, curr, end int, visited map[int]bool) bool {
    if curr == end { return true }
    visited[curr] = true
    
    for _, neighbor := range graph[curr] {
        if !visited[neighbor] {
            if dfsPath(graph, neighbor, end, visited) {
                return true
            }
        }
    }
    return false
}
```

---

### Pattern 4: Cycle Detection
**> Does this graph have a cycle?**

```go
func hasCycle(graph map[int][]int) bool {
    visited := make(map[int]bool)
    recStack := make(map[int]bool)  // Current recursion path
    
    for node := range graph {
        if !visited[node] {
            if dfsCycle(node, graph, visited, recStack) {
                return true
            }
        }
    }
    return false
}

func dfsCycle(node int, graph map[int][]int, visited, recStack map[int]bool) bool {
    visited[node] = true
    recStack[node] = true  // Add to current path
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            if dfsCycle(neighbor, graph, visited, recStack) {
                return true
            }
        } else if recStack[neighbor] {
            // Back edge found = cycle!
            return true
        }
    }
    
    recStack[node] = false  // Remove from path
    return false
}
```

**Key insight**: If we encounter a node that's in our CURRENT path (recStack), we found a cycle.

---

## 🎮 Practice Exercise

**> Problem**: Given a grid, find the size of the largest island.

<details>
<summary>Think about what to return from DFS...</summary>

Instead of void, return the size of the island from each DFS call.
</details>

<details>
<summary>Solution</summary>

```go
func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                area := dfsArea(grid, i, j)
                maxArea = max(maxArea, area)
            }
        }
    }
    return maxArea
}

func dfsArea(grid [][]int, i, j int) int {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return 0
    }
    if grid[i][j] != 1 { return 0 }
    
    grid[i][j] = 0  // Mark visited
    
    return 1 + 
        dfsArea(grid, i+1, j) +
        dfsArea(grid, i-1, j) +
        dfsArea(grid, i, j+1) +
        dfsArea(grid, i, j-1)
}
```
</details>

---

## 🔄 Iterative DFS (Avoid Stack Overflow)

```go
func dfsIterative(start int, graph map[int][]int) {
    visited := make(map[int]bool)
    stack := []int{start}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if visited[node] { continue }
        visited[node] = true
        
        // Process node...
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                stack = append(stack, neighbor)
            }
        }
    }
}
```

---

## 📊 Complexity

| Aspect | Cost | Why |
|--------|------|-----|
| Time | O(V + E) | Visit every vertex and edge once |
| Space | O(V) | Visited set + recursion stack |

V = Vertices, E = Edges

---

## ⚠️ Common Pitfalls

1. **Forgot visited check**: Infinite loop!
2. **Not marking visited before pushing**: Same node added multiple times to stack
3. **Wrong adjacency list direction**: Graph might be directed!
4. **Modifying original input**: Make copy if you need original later
5. **Stack overflow on deep graphs**: Use iterative DFS

---

## 🚀 DFS vs BFS for Graphs

| Use DFS when... | Use BFS when... |
|-----------------|-----------------|
| Exploring all paths | Finding shortest path |
| Need simple recursive solution | Need level-by-level traversal |
| Memory constrained | Has edge weights (use Dijkstra) |

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[07-tree-dfs|Tree DFS]] — Graph DFS with visited tracking (trees have no cycles)
- [[18-advanced-graphs|Advanced Graphs]] — Union Find, Topological Sort
- [[11-backtracking|Backtracking]] — DFS explores all paths
- [[23-matrix|Matrix]] — DFS on 2D grid for island problems
