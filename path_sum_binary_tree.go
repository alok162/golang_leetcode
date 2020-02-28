/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil { return false }
    sum = sum - root.Val
    if root.Left == nil && root.Right == nil && sum == 0 { return true }
    left := hasPathSum(root.Left, sum)
    right := hasPathSum(root.Right, sum)
    if left == true || right == true { 
        return true
    } else {
        return false
    }
}
