// Given an n-ary tree, return the preorder traversal of its nodes' values.

// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).


// Follow up:

// Recursive solution is trivial, could you do it iteratively?


// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,3,5,6,2,4]


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var ans []int
    if root == nil { return ans }
    ans = append(ans, root.Val)
    var preorderUtil func(root *Node)
    preorderUtil = func(root *Node) {
        for _, node := range root.Children {
            ans = append(ans, node.Val)
            preorderUtil(node)
        }
    }
    preorderUtil(root)
    return ans
}
