/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var ans [][]int
    if root == nil {return ans}
    var q []*TreeNode
    q = append(q, root)
    for len(q) > 0 {
        n := len(q)
        var tempArray []int
        for n > 0 {
            temp := q[0]
            q = q[1:]
            if temp.Left != nil {
                q = append(q, temp.Left)
            }
            if temp.Right != nil {
                q = append(q, temp.Right)
            }
            tempArray = append(tempArray, temp.Val)
            n -= 1
        }
        ans = append(ans, tempArray)
    }
    return ans
}
