package main

import "testing"

func TestLRUCache(t *testing.T) {
    lru := Constructor(2)
    lru.Put(1, 1)
    lru.Put(2, 2)
    if got := lru.Get(1); got != 1 {
        t.Errorf("Get(1) = %v, want 1", got)
    }
    lru.Put(3, 3)
    if got := lru.Get(2); got != -1 {
        t.Errorf("Get(2) = %v, want -1", got)
    }
}
