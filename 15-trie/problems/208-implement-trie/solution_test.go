package main

import "testing"

func TestTrie(t *testing.T) {
    trie := Constructor()
    trie.Insert("apple")
    if !trie.Search("apple") {
        t.Error("Search(apple) should return true")
    }
    if !trie.StartsWith("app") {
        t.Error("startsWith(app) should return true")
    }
}
