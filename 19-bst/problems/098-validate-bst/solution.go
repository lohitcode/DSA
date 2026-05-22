package main

import "math"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
    if node == nil { return true }
    if int64(node.Val) <= min || int64(node.Val) >= max { return false }
    return validate(node.Left, min, int64(node.Val)) && validate(node.Right, int64(node.Val), max)
}
