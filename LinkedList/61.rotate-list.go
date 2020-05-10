// Given a linked list, rotate the list to the right by k places, where k is non-negative.

// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    var kth_node *ListNode
    if head == nil { return head }
    // count nodes in Linked list
    n := countList(head, k)
    k = k%n
    if k == 0 { return head }
    k = n - k
    temp := head
    for temp.Next != nil {
        k --
        if k == 0 {
            kth_node = temp
        }
        temp = temp.Next
    }
    temp.Next = head
    head = kth_node.Next
    kth_node.Next = nil
    return head
}

func countList(head *ListNode, k int) int {
    temp := head
    count := 0
    for temp != nil {
        count ++
        temp = temp.Next
    }
    return count
}
