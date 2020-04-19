/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    var util func(root *TreeNode)
    sum := 0
    util = func(root *TreeNode) {
        if root == nil { return }

        util(root.Right)
        sum += root.Val
        root.Val = sum
        util(root.Left)
    }
    util(root)
    return root
}
