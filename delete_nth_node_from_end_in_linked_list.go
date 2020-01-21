/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import ("fmt")
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    constant := 0
    count := n
    var util func(head *ListNode) *ListNode
    util = func(head *ListNode) *ListNode {
        if head == nil {
            return nil
        }
        head.Next = util(head.Next)
        count --
        if count == 0 {
            return head.Next
        }
        return head
    }
    util(head)
    if constant == count {
        head = head.Next
    }
    return head
}
