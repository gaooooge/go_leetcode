package easy

import "testing"

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
    curr := head
    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }

    return prev
}

func TestReverseList(t *testing.T) {

}
