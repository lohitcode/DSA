---
title: "Pattern: Trie"
aliases: [Pattern: Trie]
type: pattern
tags:
  - pattern
---

# Trie (Prefix Tree)

## 🎯 The Core Idea

A trie is a tree where **each path from root represents a string**. Instead of storing words separately, we share common prefixes.

```
     root
     /  \
    a    b
    |    |
    p    a
   / \    |
  p   e    n
  |   |    |
  l   t    a
  |        |
  e        n
  |        |
  ✓        ✓

"apple" and "pet" share 'p'!
```

**> Quick thought**: Why would this be better than just storing words in a hash set?

<details>
<summary>Click to reveal...</summary>

Tries give you O(m) prefix searches where m = prefix length. Hash sets need O(n) to find all words with a prefix. Also, tries share memory for common prefixes.
</details>

---

## 🧠 Trie Structure

```go
type TrieNode struct {
    children [26]*TrieNode  // One for each letter (a-z)
    isEnd   bool            // Marks end of a word
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: &TrieNode{}}
}
```

**Key insight**: Each node has 26 possible children (for lowercase English). The PATH to a node spells out the prefix.

---

## 🔥 Trie Operations

### Insert: O(m) where m = word length

```go
func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        idx := ch - 'a'  // 0-25
        if node.children[idx] == nil {
            node.children[idx] = &TrieNode{}
        }
        node = node.children[idx]
    }
    node.isEnd = true  // Mark word end
}
```

**What happens**: Walk down the tree, creating nodes as needed. Mark the last node.

---

### Search: O(m)

```go
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false  // Path doesn't exist
        }
        node = node.children[idx]
    }
    return node.isEnd  // Must be a complete word
}
```

**What happens**: Follow the path. If we can traverse the whole word AND it's marked as end, found!

---

### StartsWith (Prefix Search): O(m)

```go
func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
    }
    return true  // Path exists
}
```

**What happens**: Just check if the path exists. Don't care about `isEnd`.

---

## 🔥 Advanced: With Counts

```go
type TrieNode struct {
    children [26]*TrieNode
    count    int     // How many words pass through here
    isEnd    bool
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &TrieNode{}
        }
        node = node.children[idx]
        node.count++  // Track count
    }
    node.isEnd = true
}

func (t *Trie) CountWordsWithPrefix(prefix string) int {
    node := t.root
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return 0
        }
        node = node.children[idx]
    }
    return node.count  // All words under this node
}
```

---

## 🔥 Pattern: Word Search II

**> Find all words from a dictionary in a 2D board**

```go
type TrieNode struct {
    children [26]*TrieNode
    word     string  // Store the word itself at end node
}

func findWords(board [][]byte, words []string) []string {
    // Build trie
    root := &TrieNode{}
    for _, word := range words {
        node := root
        for _, ch := range word {
            idx := ch - 'a'
            if node.children[idx] == nil {
                node.children[idx] = &TrieNode{}
            }
            node = node.children[idx]
        }
        node.word = word  // Store word at end
    }
    
    result := []string{}
    
    // DFS on board
    var dfs func(i, j int, node *TrieNode)
    dfs = func(i, j int, node *TrieNode) {
        ch := board[i][j]
        if ch == '#' || node.children[ch-'a'] == nil {
            return
        }
        
        node = node.children[ch-'a']
        if node.word != "" {
            result = append(result, node.word)
            node.word = ""  // Prevent duplicates
        }
        
        board[i][j] = '#'  // Mark visited
        if i > 0 { dfs(i-1, j, node) }
        if j > 0 { dfs(i, j-1, node) }
        if i < len(board)-1 { dfs(i+1, j, node) }
        if j < len(board[0])-1 { dfs(i, j+1, node) }
        board[i][j] = ch  // Restore
    }
    
    for i := range board {
        for j := range board[i] {
            dfs(i, j, root)
        }
    }
    
    return result
}
```

**Key insight**: The trie guides our DFS — we only explore paths that form valid prefixes!

---

## 🎮 Practice Exercise

**> Problem**: Design a data structure that supports adding new words and finding if a string matches any previously added string OR has exactly one dot wildcard (.)

Example: `.ad` matches `"bad"`, `"dad"`, `"mad"` but not `"pad"`

<details>
<summary>Hint: How do you handle the wildcard?</summary>

When you see `.`, try ALL 26 possible children!
</details>

<details>
<summary>Solution</summary>

```go
type WordDictionary struct {
    root *TrieNode
}

func (wd *WordDictionary) AddWord(word string) {
    node := wd.root
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &TrieNode{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
    return wd.search(word, 0, wd.root)
}

func (wd *WordDictionary) search(word string, i int, node *TrieNode) bool {
    if i == len(word) {
        return node.isEnd
    }
    
    ch := word[i]
    if ch == '.' {
        // Try all children
        for _, child := range node.children {
            if child != nil && wd.search(word, i+1, child) {
                return true
            }
        }
        return false
    }
    
    idx := ch - 'a'
    if node.children[idx] == nil {
        return false
    }
    return wd.search(word, i+1, node.children[idx])
}
```
</details>

---

## 📊 Complexity

| Operation | Time | Space |
|-----------|------|-------|
| Insert | O(m) | O(m) |
| Search | O(m) | O(1) |
| StartsWith | O(m) | O(1) |
| Delete | O(m) | O(1) |

m = length of word

---

## ⚠️ Common Pitfalls

1. **Fixed alphabet**: Using `[26]*TrieNode` assumes lowercase English
2. **Memory overhead**: Tries use more memory than hash sets for small dictionaries
3. **Case sensitivity**: `A` vs `a` are different unless you normalize
4. **Not clearing `isEnd`**: For search/replace, might need to unset

---

## 🚀 When to Use a Trie

✅ **Use Trie when:**
- Need prefix-based searches (autocomplete, spell check)
- Many strings share common prefixes
- Need to find all words with a given prefix
- Building a dictionary

❌ **Don't use Trie when:**
- Few strings with little common prefix
- Only need exact word lookup (hash set is simpler)
- Memory is constrained
- Long unique strings (little sharing)

---

## 💡 Trie Variants

| Variant | Use Case |
|---------|----------|
| Suffix Trie | Pattern matching in text |
| Radix Tree | Compressed trie (merge chains) |
| Ternary Search Tree | 3 children per node (less memory) |
| B-Tree | Disk-based storage (databases) |

---

## 📚 Real-World Applications

- **Autocomplete**: Google search, IDE suggestions
- **Spell check**: Find words with given prefix
- **IP routing**: Longest prefix match
- **T9 predictive text**: Phone number to words
- **Word games**: Scrabble, Boggle solver

---
**[[../index.md|← Back to Topics]]**

## Related Topics
- [[03-arrays-hashing|Arrays & Hashing]] — Trie vs hash map for string operations
- [[22-string-advanced|String Advanced]] — Trie enables prefix-based string operations
- [[21-design/21-design-problems|Design Problems]] — Trie is used in autocomplete design
