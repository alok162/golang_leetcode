/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var util func(root *TreeNode) 
    
    var prev *TreeNode
    flag := true
    util = func(root *TreeNode) {
        if root == nil {return}
        util(root.Left)
        if prev != nil && prev.Val >= root.Val {
                flag = false
        }
        prev = root
        util(root.Right)
    }
    util(root)
    return flag
}
