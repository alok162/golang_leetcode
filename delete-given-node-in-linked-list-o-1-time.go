/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    temp := node
    temp = temp.Next
    temp.Val, node.Val = node.Val, temp.Val
    node.Next = temp.Next
    temp = nil
}
