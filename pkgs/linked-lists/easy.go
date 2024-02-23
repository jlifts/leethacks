package linkedlists

// https://leetcode.com/problems/merge-two-sorted-lists/
type ListNode struct {
      Val int
      Next *ListNode
 }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    empty := &ListNode{}
    tail := empty

    for list1 != nil && list2 != nil {
       if list1.Val < list2.Val {
           tail.Next = list1
           list1 = list1.Next
       } else {
           tail.Next = list2
           list2 = list2.Next
       }

       tail = tail.Next
    }

    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }

    return empty.Next
}

// https://leetcode.com/problems/reverse-linked-list/description/

// Iterative T: O(n) M: O(1)
func reverseList(head *ListNode) *ListNode {
    curr :=  head
    var prev *ListNode
    for curr != nil {
        nxt := curr.Next
        curr.Next = prev
        prev = curr
        curr = nxt
    }
    return prev
}
// T&M: O(n)
// Recursive
 func recursiveReverseList(head *ListNode) *ListNode {
     if head == nil {
         return nil
     }

     newHead := head
     if head.Next != nil {
         newHead = reverseList(head.Next)
         head.Next.Next = head
     }
     head.Next = nil

     return newHead
 }
