/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

 if head == nil {
    return head
 }   

curr := head

curr2 := head

count := 0

count2 := 0

for curr.Next != nil {
count ++
curr = curr.Next
}

index := count - n+1


var prev = &ListNode{}

if index == 0 {
    head =  head.Next
    return head
}

for curr2 != nil {
    if count2 == index {
        prev.Next = curr2.Next
    }
    prev = curr2
    curr2 = curr2.Next
    count2++
}
return head
}