package main

type DListNode struct {
    key, val   int
    prev, next *DListNode
}

type LRUCache struct {
    capacity int
    cache    map[int]*DListNode
    head, tail *DListNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{capacity: capacity, cache: make(map[int]*DListNode)}
    lru.head = &DListNode{}
    lru.tail = &DListNode{}
    lru.head.next = lru.tail
    lru.tail.prev = lru.head
    return lru
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        this.add(node)
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        node.val = value
        this.add(node)
    } else {
        node := &DListNode{key: key, val: value}
        this.cache[key] = node
        this.add(node)
        if len(this.cache) > this.capacity {
            delete(this.cache, this.tail.prev.key)
            this.remove(this.tail.prev)
        }
    }
}

func (this *LRUCache) remove(node *DListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) add(node *DListNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}
