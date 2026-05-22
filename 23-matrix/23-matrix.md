---
title: "Pattern: Matrix"
aliases: [Pattern: Matrix]
type: pattern
tags:
  - pattern
---

# Matrix (2D Arrays)

## 🎯 The Core Idea

A matrix is just a **2D grid** of elements. What makes matrix problems interesting is how you **traverse** and **manipulate** the grid — often treating it as a graph or applying spatial transformations.

**> Quick thought**: Why are matrices so common in coding problems?</summary>

<details>
<summary>Click to reveal...</summary>

Matrices represent images, game boards, spreadsheets, graphs, and more. They test your ability to think in 2D and handle edge cases (boundaries!).
</details>

---

## 🧠 Matrix Fundamentals

```go
// Declaration
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Dimensions
m := len(matrix)    // Rows
n := len(matrix[0]) // Cols (assume non-empty)

// Access
element := matrix[i][j]  // Row i, Column j

// Iterate all elements
for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
        // Process matrix[i][j]
    }
}
```

**Key insight**: `matrix[i]` is the i-th ROW. `matrix[i][j]` is element at row i, column j.

---

## 🔥 Essential Matrix Patterns

### Pattern 1: Spiral Traversal
**> Visit matrix in spiral order (layer by layer)**

```go
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    m, n := len(matrix), len(matrix[0])
    top, bottom := 0, m-1
    left, right := 0, n-1
    result := []int{}
    
    for top <= bottom && left <= right {
        // Traverse right (top row)
        for j := left; j <= right; j++ {
            result = append(result, matrix[top][j])
        }
        top++
        
        // Traverse down (right column)
        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--
        
        // Traverse left (bottom row) — if still valid
        if top <= bottom {
            for j := right; j >= left; j-- {
                result = append(result, matrix[bottom][j])
            }
            bottom--
        }
        
        // Traverse up (left column) — if still valid
        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }
    
    return result
}
```

**Visual**:
```
→ → → ↓
↑   → ↓
↑ ← ← ↓
↑ ← ← ←
```

**Key insight**: Maintain four boundaries and shrink them after processing each side.

---

### Pattern 2: In-Place Rotation (90° Clockwise)
**> Rotate matrix without using extra space**

```go
func rotate(matrix [][]int) {
    n := len(matrix)
    
    // Step 1: Transpose (swap across diagonal)
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    // Step 2: Reverse each row
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}
```

**Visual**:
```
Original:    Transpose:   Reverse rows:
1 2 3        1 4 7        7 4 1
4 5 6   →    2 5 8   →    8 5 2
7 8 9        3 6 9        9 6 3
```

**Why it works**: Transpose converts rows to columns, reversing completes the rotation.

---

### Pattern 3: Set Matrix Zeroes (In-Place)
**> If element is 0, set its entire row/col to 0**

```go
func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])
    
    // Use first row/col as markers
    firstRowZero, firstColZero := false, false
    
    // Check if first row/col needs zeroing
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            firstRowZero = true
            break
        }
    }
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            firstColZero = true
            break
        }
    }
    
    // Use first row/col to mark zeros
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    
    // Zero out based on markers
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    
    // Zero first row/col if needed
    if firstRowZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if firstColZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
```

**Key insight**: Use O(1) extra space by repurposing first row/col as markers.

---

### Pattern 4: Search in Sorted Matrix
**> Matrix where each row/col is sorted**

```go
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    
    // Start from top-right corner
    row, col := 0, len(matrix[0])-1
    
    for row < len(matrix) && col >= 0 {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            col--  // Move left (smaller values)
        } else {
            row++  // Move down (larger values)
        }
    }
    
    return false
}
```

**Why start at top-right?**:
- Left is smaller, down is larger
- Can eliminate row or column each step
- O(m + n) time instead of O(m × n)

---

### Pattern 5: DFS/BFS in Matrix
**> Treat matrix as graph, explore connected regions**

```go
// Island count (connected 1s)
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    count := 0
    m, n := len(grid), len(grid[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                count++
            }
        }
    }
    
    return count
}

func dfs(grid [][]byte, i, j int) {
    m, n := len(grid), len(grid[0])
    
    // Boundary / water check
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
        return
    }
    
    grid[i][j] = '0'  // Mark visited
    
    // Explore 4 directions
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j-1)
}
```

**Key insight**: Modify matrix in-place to mark visited (saves memory).

---

## 🎮 Practice Exercise

**> Problem: Given an image represented as matrix, flip it horizontally, then invert.**

**Flip**: Reverse each row. **Invert**: 0→1, 1→0.

<details>
<summary>Think about how to do this in one pass...</summary>

For each row, iterate from both ends toward center, swapping and inverting.
</details>

<details>
<summary>Solution</summary>

```go
func flipAndInvertImage(image [][]int) [][]int {
    n := len(image)
    
    for i := 0; i < n; i++ {
        left, right := 0, n-1
        for left <= right {
            // Swap
            image[i][left], image[i][right] = image[i][right], image[i][left]
            
            // Invert (only need to invert one side since we swapped)
            if left != right {  // Don't double-invert middle
                image[i][left] ^= 1
                image[i][right] ^= 1
            } else {
                image[i][left] ^= 1
            }
            
            left++
            right--
        }
    }
    
    return image
}
```
</details>

---

## 📊 Matrix Operations Complexity

| Operation | Time | Space |
|-----------|------|-------|
| Traverse all | O(m × n) | O(1) |
| Spiral | O(m × n) | O(1) output |
| Rotate in-place | O(m × n) | O(1) |
| Set zeroes | O(m × n) | O(1) |
| Sorted search | O(m + n) | O(1) |
| DFS/BFS | O(m × n) | O(m × n) worst |

---

## ⚠️ Common Pitfalls

1. **Boundary checks**: Always validate `i >= 0 && i < m && j >= 0 && j < n`
2. **Empty matrix**: Check `len(matrix) == 0` before accessing `matrix[0]`
3. **Off-by-one**: Remember it's `matrix[m-1][n-1]` not `matrix[m][n]`
4. **Jagged arrays**: Not all rows have same length (usually not the case in problems)
5. **Modifying during iteration**: Be careful when changing values while iterating

---

## 🚀 Direction Patterns

```go
// 4-directional (up, down, left, right)
dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 8-directional (includes diagonals)
dirs := [][]int{
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1},          {0, 1},
    {1, -1},  {1, 0},  {1, 1},
}

// Usage
for _, dir := range dirs {
    newRow, newCol := i + dir[0], j + dir[1]
    // Check bounds and process
}
```

---

## 💡 Matrix as Graph

- Each cell = node
- Adjacent cells = edges (4 or 8 directional)
- Use DFS/BFS for island problems, path finding
- Use Dijkstra for weighted grid problems

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[16-dp-2d|DP 2D]] — Many DP problems use matrix
- [[10-graphs-dfs|Graph DFS]] — DFS on matrix for island problems
- [[20-math-geometry|Math & Geometry]] — Matrix rotations and transformations
