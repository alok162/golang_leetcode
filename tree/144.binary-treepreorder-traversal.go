// Given a binary tree, return the preorder traversal of its nodes' values.

// Example:

// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [1,2,3]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var ans []int
    preorderTraversalUtil(root, &ans)
    return ans
}

func preorderTraversalUtil(root *TreeNode, ans *[]int) {
    if root == nil {
        return
    }
    *ans = append(*ans, root.Val)
    preorderTraversalUtil(root.Left, ans)
    preorderTraversalUtil(root.Right, ans)
}
