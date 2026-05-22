---
title: "Pattern: Advanced Graphs"
aliases: [Pattern: Advanced Graphs]
type: pattern
tags:
  - pattern
---

# Advanced Graphs

## 🎯 The Core Idea

Basic DFS/BFS explores graphs. **Advanced graph algorithms** solve harder problems: shortest paths, minimum spanning trees, detecting cycles, and processing dependencies.

**> Quick thought**: Why can't we just use BFS for everything?

<details>
<summary>Click to reveal...</summary>

BFS finds shortest path in **unweighted** graphs. For weighted graphs, negative edges, or finding minimum cost to connect all nodes, we need specialized algorithms.
</details>

---

## 🧠 The Advanced Graph Toolkit

| Algorithm | What it does | When to use |
|-----------|--------------|-------------|
| Dijkstra | Shortest path (positive weights) | Weighted graphs, navigation |
| Bellman-Ford | Shortest path (negative OK) | Detect negative cycles |
| Floyd-Warshall | All-pairs shortest path | Dense graphs, small n |
| Topological Sort | Linear ordering of dependencies | Course scheduling, build systems |
| Union Find | Connected components, cycle detection | Dynamic connectivity |
| Prim's/Kruskal's | Minimum spanning tree | Network design, clustering |
| A* | Shortest path with heuristic | Games, GPS, pathfinding |

---

## 🔥 Essential Algorithms

### Algorithm 1: Dijkstra's Shortest Path
**> Find shortest path from source to all nodes (non-negative weights)**

```go
type Edge struct {
    to, weight int
}

func dijkstra(start int, graph [][]Edge, n int) []int {
    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt32
    }
    dist[start] = 0

    // Min-heap: [distance, node]
    pq := &MinHeap{}
    heap.Init(pq)
    heap.Push(pq, [2]int{0, start})

    for pq.Len() > 0 {
        curr := heap.Pop(pq).([2]int)
        d, node := curr[0], curr[1]

        // Skip outdated entries
        if d > dist[node] {
            continue
        }

        for _, edge := range graph[node] {
            newDist := dist[node] + edge.weight
            if newDist < dist[edge.to] {
                dist[edge.to] = newDist
                heap.Push(pq, [2]int{newDist, edge.to})
            }
        }
    }

    return dist
}
```

**Key insight**: Always process the unvisited node with smallest distance first. Once a node is "settled", we've found its shortest distance.

**Why heap?**: O((V+E) log V) vs O(V²) naive implementation.

---

### Algorithm 2: Union Find (Disjoint Set)
**> Track connected components with path compression and union by rank**

```go
type UnionFind struct {
    parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make([]int, n),
        rank:   make([]int, n),
    }
    for i := range uf.parent {
        uf.parent[i] = i
    }
    return uf
}

func (uf *UnionFind) Find(x int) int {
    // Path compression: point directly to root
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
    px, py := uf.Find(x), uf.Find(y)
    if px == py {
        return false  // Already connected
    }

    // Union by rank: attach smaller tree to larger
    if uf.rank[px] < uf.rank[py] {
        uf.parent[px] = py
    } else if uf.rank[px] > uf.rank[py] {
        uf.parent[py] = px
    } else {
        uf.parent[py] = px
        uf.rank[px]++
    }
    return true
}

func (uf *UnionFind) Connected(x, y int) bool {
    return uf.Find(x) == uf.Find(y)
}
```

**Key optimizations**:
- **Path compression**: `Find` flattens the tree
- **Union by rank**: Keeps trees shallow

**Result**: Nearly O(1) amortized for both operations!

---

### Algorithm 3: Topological Sort
**> Linear ordering where dependencies come before dependents**

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    // Build graph and count in-degrees
    graph := make(map[int][]int)
    inDegree := make([]int, numCourses)

    for _, pre := range prerequisites {
        course, prereq := pre[0], pre[1]
        graph[prereq] = append(graph[prereq], course)
        inDegree[course]++
    }

    // Start with courses that have no prerequisites
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    completed := 0
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        completed++

        // "Remove" this course — reduce in-degrees
        for _, next := range graph[course] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    return completed == numCourses
}
```

**Kahn's Algorithm**: Repeatedly remove nodes with no incoming edges.

**Cycle detection**: If topological sort has < n nodes, there's a cycle!

---

### Algorithm 4: Minimum Spanning Tree (Prim's)
**> Connect all nodes with minimum total edge weight**

```go
type Edge struct {
    to, weight int
}

func primMST(n int, edges [][]int) int {
    // Build adjacency list
    graph := make([][]Edge, n)
    for _, e := range edges {
        from, to, weight := e[0], e[1], e[2]
        graph[from] = append(graph[from], Edge{to, weight})
        graph[to] = append(graph[to], Edge{from, weight})
    }

    visited := make([]bool, n)
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    // Start from node 0
    visited[0] = true
    for _, e := range graph[0] {
        heap.Push(minHeap, [2]int{e.weight, e.to})
    }

    totalWeight := 0
    edgesUsed := 0

    for minHeap.Len() > 0 && edgesUsed < n-1 {
        curr := heap.Pop(minHeap).([2]int)
        weight, to := curr[0], curr[1]

        if visited[to] {
            continue
        }

        visited[to] = true
        totalWeight += weight
        edgesUsed++

        for _, e := range graph[to] {
            if !visited[e.to] {
                heap.Push(minHeap, [2]int{e.weight, e.to})
            }
        }
    }

    if edgesUsed != n-1 {
        return -1  // Not connected
    }
    return totalWeight
}
```

**Greedy choice**: Always add the cheapest edge that connects a new node.

---

## 🎮 Practice Exercise

**> Problem: Number of Connected Components in an Undirected Graph**

<details>
<summary>Think about which tool to use...</summary>

Union Find is perfect for this! Or DFS/BFS to count components.
</details>

<details>
<summary>Union Find Solution</summary>

```go
func countComponents(n int, edges [][]int) int {
    uf := NewUnionFind(n)
    components := n

    for _, edge := range edges {
        if uf.Union(edge[0], edge[1]) {
            components--
        }
    }

    return components
}
```
</details>

---

## 📊 Complexity Comparison

| Algorithm | Time | Space |
|-----------|------|-------|
| Dijkstra (heap) | O((V+E) log V) | O(V+E) |
| Dijkstra (array) | O(V²) | O(V+E) |
| Bellman-Ford | O(VE) | O(V+E) |
| Floyd-Warshall | O(V³) | O(V²) |
| Topological Sort | O(V+E) | O(V+E) |
| Union Find | ~O(α(n)) per op | O(V) |
| Prim's MST | O((V+E) log V) | O(V+E) |
| Kruskal's MST | O(E log E) | O(V+E) |

V = vertices, E = edges, α(n) ≈ 1 (inverse Ackermann)

---

## ⚠️ Common Pitfalls

1. **Negative weights**: Dijkstra FAILS with negative weights (use Bellman-Ford)
2. **Disconnected graphs**: Topological sort might not include all nodes
3. **Cycle detection**: Union Find tells you IF cycle exists, not where
4. **Heap duplicates**: Dijkstra may push same node multiple times (check distance)
5. **Integer overflow**: Path sums can exceed int range

---

## 🚀 When to Use Each

### Use Dijkstra when:
- Non-negative edge weights
- Single source shortest path
- Navigation, routing

### Use Bellman-Ford when:
- Negative edge weights possible
- Need to detect negative cycles
- Small graphs

### Use Floyd-Warshall when:
- Need all-pairs shortest paths
- Dense graph (E ≈ V²)
- Small graph (n < 400)

### Use Topological Sort when:
- Processing dependencies
- Course scheduling
- Build order determination

### Use Union Find when:
- Dynamic connectivity
- Kruskal's MST
- Cycle detection in undirected graphs

---

## 💡 Real-World Applications

| Algorithm | Application |
|-----------|-------------|
| Dijkstra | GPS navigation, network routing |
| A* | Games, pathfinding with heuristics |
| Topological Sort | Task scheduling, package managers |
| MST | Network design, clustering |
| Union Find | Social networks (friend groups), image segmentation |
| Bellman-Ford | Financial arbitrage detection |

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[10-graphs-dfs|Graph DFS]] — Foundation for graph algorithms
- [[12-heap|Heap]] — Dijkstra uses min-heap
- [[09-tree-bfs|Tree BFS]] — BFS is used in level-order and shortest path
