/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"
func maxLevelSum(root *TreeNode) int {
    if root == nil {return 0}
    var q []*TreeNode
    q = append(q, root)
    result := 0
    level := 0
    curr_level := 0
    for len(q) > 0 {
        count := len(q)
        curr_level += 1
        sum := 0
        for count > 0 {
            temp := q[0]
            count -= 1
            q = q[1:]
            sum += temp.Val
            if temp.Left != nil {
                q = append(q, temp.Left)
            }
            if temp.Right != nil {
                q = append(q, temp.Right)
            }
            if count == 0 {
                if result < sum {
                    result = sum
                    level = curr_level
                }
            }
        }
    }
    return level
}
