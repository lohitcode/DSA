package main

type LFUCache struct {
    capacity   int
    minFreq    int
    keyToVal   map[int]int
    keyToFreq  map[int]int
    freqToKeys map[int]map[int]bool
}

func Constructor460(capacity int) LFUCache {
    return LFUCache{
        capacity:   capacity,
        keyToVal:   make(map[int]int),
        keyToFreq:  make(map[int]int),
        freqToKeys: make(map[int]map[int]bool),
    }
}

func (this *LFUCache) Get(key int) int {
    if val, ok := this.keyToVal[key]; ok {
        this.updateFreq(key)
        return val
    }
    return -1
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 { return }
    
    if _, ok := this.keyToVal[key]; ok {
        this.keyToVal[key] = value
        this.updateFreq(key)
    } else {
        if len(this.keyToVal) >= this.capacity {
            this.evict()
        }
        this.keyToVal[key] = value
        this.keyToFreq[key] = 1
        this.minFreq = 1
        if this.freqToKeys[1] == nil {
            this.freqToKeys[1] = make(map[int]bool)
        }
        this.freqToKeys[1][key] = true
    }
}

func (this *LFUCache) updateFreq(key int) {
    freq := this.keyToFreq[key]
    this.keyToFreq[key] = freq + 1
    delete(this.freqToKeys[freq], key)
    if this.freqToKeys[freq+1] == nil {
        this.freqToKeys[freq+1] = make(map[int]bool)
    }
    this.freqToKeys[freq+1][key] = true
    if freq == this.minFreq && len(this.freqToKeys[freq]) == 0 {
        this.minFreq++
    }
}

func (this *LFUCache) evict() {
    for key := range this.freqToKeys[this.minFreq] {
        delete(this.keyToVal, key)
        delete(this.keyToFreq, key)
        delete(this.freqToKeys[this.minFreq], key)
        return
    }
}
