/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    var q []*Node
    var ans [][]int
    if root == nil { return ans}
    q = append(q, root)
    for len(q) > 0 {
        n := len(q)
        var tempArray []int
        for n > 0 {
            temp := q[0]
            q = q[1:]
            tempArray = append(tempArray, temp.Val)
            
            for _, node := range temp.Children {
                if node != nil {
                    q = append(q, node)
                }
            }
            n -= 1
        }
        ans = append(ans, tempArray)
    }
   
    return ans
}
