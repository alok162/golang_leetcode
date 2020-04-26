/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given a binary tree
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

func levelOrderBottom(root *TreeNode) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    var q []*TreeNode
    q = append(q, root)
    for len(q) > 0 {
        levelSize := len(q)
        var tempArr []int
        for levelSize > 0 {
            temp := q[0]
            q = q[1:]
            tempArr = append(tempArr, temp.Val)
            if temp.Left != nil {
                q = append(q, temp.Left)
            }
            if temp.Right != nil {
                q = append(q, temp.Right)
            }
            levelSize -= 1
        }
        res = append(res, tempArr)
    }
    return reverse(res)
}

func reverse(res [][]int) [][]int {
    i, j := 0, len(res)-1
    for i <= j {
        res[i], res[j] = res[j], res[i]
        i += 1
        j -= 1
    }
    return res
}
