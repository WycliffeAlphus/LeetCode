/**
 * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
dummyHead := &ListNode{}
curr := dummyHead
carry := 0
var toadd int
for l1 != nil || l2 != nil || carry > 0 {
    val1, val2 := 0, 0
    if l1 != nil{
        val1 = l1.Val
        l1 = l1.Next
    }

    if l2 != nil {
        val2 = l2.Val
        l2 = l2.Next
    }

    value := val1 + val2 + carry
    toadd = value%10 
    curr.Next = &ListNode{Val:toadd}
    carry = value/10
    curr = curr.Next
}

return dummyHead.Next

}