package main

import (
    "container/heap"
    "math"
)

type Edge struct{ to, weight int }

func networkDelayTime(times [][]int, n int, k int) int {
    adj := make([][]Edge, n+1)
    for _, t := range times {
        adj[t[0]] = append(adj[t[0]], Edge{t[1], t[2]})
    }
    
    dist := make([]int, n+1)
    for i := range dist { dist[i] = math.MaxInt32 }
    dist[k] = 0
    
    mh := &MinHeap{[2]int{0, k}}
    heap.Init(mh)
    
    for mh.Len() > 0 {
        curr := heap.Pop(mh).([2]int)
        d, node := curr[0], curr[1]
        if d > dist[node] { continue }
        for _, e := range adj[node] {
            if d+e.weight < dist[e.to] {
                dist[e.to] = d + e.weight
                heap.Push(mh, [2]int{dist[e.to], e.to})
            }
        }
    }
    
    maxDist := 0
    for i := 1; i <= n; i++ {
        if dist[i] == math.MaxInt32 { return -1 }
        if dist[i] > maxDist { maxDist = dist[i] }
    }
    return maxDist
}

type MinHeap [][2]int // [distance, node]

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old) - 1
    x := old[n]
    *h = old[:n]
    return x
}
