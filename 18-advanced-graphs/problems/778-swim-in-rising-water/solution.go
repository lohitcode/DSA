package main

import "container/heap"

func swimInWater(grid [][]int) int {
    n := len(grid)
    dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    
    visited := make([][]bool, n)
    for i := range visited { visited[i] = make([]bool, n) }
    
    mh := &MinHeap{[3]int{grid[0][0], 0, 0}}
    heap.Init(mh)
    
    maxTime := 0
    for mh.Len() > 0 {
        curr := heap.Pop(mh).([3]int)
        t, r, c := curr[0], curr[1], curr[2]
        maxTime = max(maxTime, t)
        if r == n-1 && c == n-1 { return maxTime }
        if visited[r][c] { continue }
        visited[r][c] = true
        
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
                heap.Push(mh, [3]int{grid[nr][nc], nr, nc})
            }
        }
    }
    return -1
}

func max(a, b int) int { if a > b { return a }; return b }

type MinHeap [][3]int // [time, row, col]

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old) - 1
    x := old[n]
    *h = old[:n]
    return x
}
