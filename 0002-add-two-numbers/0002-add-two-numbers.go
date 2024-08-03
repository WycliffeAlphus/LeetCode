/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

carry := 0
result := &ListNode{}
 current := result 
for l1 != nil || l2 != nil || carry != 0 {
  var  val1, val2 int
if l1 != nil {
val1 = l1.Val
l1 = l1.Next
}
    if l2 != nil {
val2 = l2.Val
l2 = l2.Next
    }   
        
        sum := val1 + val2 + carry
        current.Next = &ListNode{Val: sum%10}
        carry = sum/10
        current = current.Next
}  

return result.Next

}