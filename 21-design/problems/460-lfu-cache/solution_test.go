package main

import "testing"

func TestLFUCache(t *testing.T) {
    lfu := Constructor460(2)
    lfu.Put(1, 1)
    lfu.Put(2, 2)
    if got := lfu.Get(1); got != 1 {
        t.Errorf("Get(1) = %v, want 1", got)
    }
    lfu.Put(3, 3)
    if got := lfu.Get(2); got != -2 {
        t.Errorf("Get(2) should be -1 (evicted), got %v", got)
    }
}
