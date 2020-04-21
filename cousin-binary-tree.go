/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var q []*TreeNode
    var parentX *TreeNode
    var parentY *TreeNode
    var labelX int
    var labelY int
    label := 0
    q = append(q, root)
    for len(q) > 0 {
        n := len(q)
        label += 1
        for n > 0 {
            temp := q[0]
            q = q[1:]
            if temp.Left != nil {
                q = append(q, temp.Left)
                if temp.Left.Val == x {
                    parentX = temp
                    labelX = label + 1
                }
                 if temp.Left.Val == y {
                    parentY = temp
                    labelY = label + 1
                }
            }
            if temp.Right != nil {
                q = append(q, temp.Right)
                if temp.Right.Val == y {
                    parentY = temp
                    labelY = label + 1
                }
                if temp.Right.Val == x {
                    parentX = temp
                    labelX = label + 1
                }
            }
            n -= 1
        }
        if parentX != nil && parentY != nil {
            break
        }
    }
    if labelX == labelY && parentX != parentY {
        return true
    } else {
        return false
    }
}
