/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil { return true}
    return isSymmetricUtil(root.Left, root.Right)
}

func isSymmetricUtil(root1 *TreeNode, root2 *TreeNode) bool {
    
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 != nil && root2 != nil {
        if root1.Val == root2.Val && isSymmetricUtil(root1.Left, root2.Right) &&
            isSymmetricUtil(root1.Right, root2.Left) {
                return true
            } else {
                return false
            }
    } else {
        return false
    }
}
