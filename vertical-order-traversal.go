/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
    node *TreeNode
    ht int
}

func verticalTraversal(root *TreeNode) [][]int {
    var keys []int
    var x [][]int
    if root == nil {return x}
    var q []Node
    ans := make(map[int][]int)
    q = append(q, Node{node: root, ht: 0})
    
    for len(q) > 0 {
        n := len(q)
        for n > 0 {
            temp := q[0]
            q = q[1:]
            ans[temp.ht] = append(ans[temp.ht], temp.node.Val)
            if temp.node.Left != nil {
                q = append(q, Node{node: temp.node.Left, ht: temp.ht-1})
            }
            if temp.node.Right != nil {
                q = append(q, Node{node: temp.node.Right, ht: temp.ht+1})
            }
            n -= 1
        }
    }
    for key, _ := range ans {
        keys = append(keys, key)
    }
    sort.Ints(keys)
    for _, val := range keys {
        x = append(x, ans[val])
    }
    return x
}
