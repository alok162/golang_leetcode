/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"
func zigzagLevelOrder(root *TreeNode) [][]int {
    var leftStack []*TreeNode
    var rightStack []*TreeNode
    var result [][]int
    if root == nil {return nil}
    leftStack = append(leftStack, root)
    flag := true
    for flag == true {
        
        // iterate left stack         
        var tempArray1 []int
        for len(leftStack) > 0 {
            temp := leftStack[len(leftStack)-1]
            leftStack = leftStack[:len(leftStack)-1]
            tempArray1 = append(tempArray1, temp.Val)
            if temp.Left != nil {
                rightStack = append(rightStack, temp.Left)
            }
            if temp.Right != nil {
                rightStack = append(rightStack, temp.Right)
            }
        }
        if len(tempArray1) > 0 {
            result = append(result, tempArray1)
        }
        
        // iterate right stack
        var tempArray2 []int
        for len(rightStack) > 0 {
            temp := rightStack[len(rightStack)-1]
            rightStack = rightStack[:len(rightStack)-1]
            tempArray2 = append(tempArray2, temp.Val)
            if temp.Right != nil {
                leftStack = append(leftStack, temp.Right)
            }
            if temp.Left != nil {
                leftStack = append(leftStack, temp.Left)
            }
        }
        if len(tempArray2) > 0 {
            result = append(result, tempArray2)
        }
        if len(leftStack) == 0 && len(rightStack) == 0 {
            flag = false
        }
    }
    return result
}
