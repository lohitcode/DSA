package main

import (
    "container/heap"
    "math"
)

func minCostConnectPoints(points [][]int) int {
    n := len(points)
    visited := make([]bool, n)
    minHeap := &MinHeap{}
    heap.Init(minHeap)
    heap.Push(minHeap, [3]int{0, 0, 0}) // [cost, from, to]
    
    totalCost := 0
    edgesUsed := 0
    
    for edgesUsed < n {
        curr := heap.Pop(minHeap).([3]int)
        cost, from, to := curr[0], curr[1], curr[2]
        if visited[to] { continue }
        visited[to] = true
        totalCost += cost
        edgesUsed++
        
        for i := 0; i < n; i++ {
            if !visited[i] {
                dist := manhattan(points[to], points[i])
                heap.Push(minHeap, [3]int{dist, to, i})
            }
        }
    }
    return totalCost
}

func manhattan(p1, p2 []int) int {
    return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int { if x < 0 { return -x }; return x }

type MinHeap [][3]int

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
