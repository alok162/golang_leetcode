// Given a binary tree, find the leftmost value in the last row of the tree.

// Example 1:
// Input:

//     2
//    / \
//   1   3

// Output:
// 1
// Example 2:
// Input:

//         1
//        / \
//       2   3
//      /   / \
//     4   5   6
//        /
//       7

// Output:
// 7


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func findBottomLeftValue(root *TreeNode) int {
    var ans int
    var q []*TreeNode
    q = append(q, root)
    
    for len(q) > 0 {
        n := len(q)
        flag := false
        
        for n > 0 {
            temp := q[0]
            q = q[1:]
            if flag == false {
                ans = temp.Val
                flag = true
            }
            if temp.Left != nil {
                q = append(q, temp.Left)
            }
            if temp.Right != nil {
                q = append(q, temp.Right)
            }
            n --
        }
    }
    return ans
}
