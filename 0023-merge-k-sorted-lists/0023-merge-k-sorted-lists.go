/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    
dummy := &ListNode{}
tail := dummy
for _, ls := range lists {
    for ls != nil {
        next := ls.Next
        tail.Next = ls
    tail = tail.Next
    ls = next 
    }
}
return sortlist(dummy.Next)
}


func sortlist(l *ListNode) *ListNode{

    curr := l

    if l == nil {
        return nil
    }

    for curr != nil {
        index := curr.Next

        for index != nil {
            if curr.Val > index.Val {
                temp := curr.Val
                curr.Val = index.Val
                index.Val = temp
            }

            index = index.Next
        }

        curr = curr.Next

    }
    return l
}