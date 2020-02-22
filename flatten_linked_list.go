/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    last := root
    var util func (root *TreeNode)
    util = func(root *TreeNode)  {
        if root == nil { return }
        left := root.Left
        right := root.Right
        if last != root {
            last.Right = root
            last.Left = nil
            last = root
        }
        util(left)
        util(right)
        if root.Left == nil && root.Right == nil { last = root }
    }
    util(root)
}
