/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int
    
    var util func(root *TreeNode)
    
    util = func(root *TreeNode) {
        if root == nil {
            return
        }
        util(root.Left)
        res = append(res, root.Val)
        util(root.Right)
    }
    util(root)
    return res
    
}
