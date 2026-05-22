package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 { return nil }
    
    rootVal := preorder[0]
    root := &TreeNode{Val: rootVal}
    
    splitIdx := 0
    for i, val := range inorder {
        if val == rootVal {
            splitIdx = i
            break
        }
    }
    
    root.Left = buildTree(preorder[1:1+splitIdx], inorder[:splitIdx])
    root.Right = buildTree(preorder[1+splitIdx:], inorder[splitIdx+1:])
    
    return root
}
