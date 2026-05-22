package main

type Trie struct {
    children [26]*TrieNode
    isEnd    bool
}

func Constructor() Trie {
    return Trie{}
}

func (t *Trie) Insert(word string) {}
func (t *Trie) Search(word string) bool {
    return false
}
func (t *Trie) StartsWith(prefix string) bool {
    return false
}
