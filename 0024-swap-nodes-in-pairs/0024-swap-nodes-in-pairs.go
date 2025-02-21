/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

    if head == nil || head.Next == nil {
        return head
    }

    curr := head

    newHead := head.Next

     var prev *ListNode


    for curr != nil &&  curr.Next != nil {
        next := curr.Next
        
        nextNext := next.Next

        curr.Next.Next = curr

        curr.Next = nextNext

        if prev != nil {
            prev.Next = next
        }
        prev = curr
        curr = nextNext

    }

    return newHead

}