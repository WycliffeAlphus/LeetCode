/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    curr := head
    var prev *ListNode

    for curr != nil {
       
if prev != nil && prev.Val == curr.Val {
            prev.Next = curr.Next
            curr = curr.Next
        } else {
            prev = curr
            curr = curr.Next
        }
        
    }
    return head
}