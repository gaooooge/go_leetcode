package easy

import (
	"fmt"
	"testing"
)

func hasCycle(head *ListNode) bool {
	hash := map[*ListNode]bool{}

	for head != nil {
		if ok := hash[head]; ok {
			return true
		}
		head = head.Next
		hash[head] = true
	}

	return false
}

func SetList(nums []int) *ListNode {
	var prev *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		var a *ListNode
		a.Val = nums[i]
		a.Next = prev
		prev = a
	}
	return prev
}

func TestHasCycle(t *testing.T) {
	l := SetList([]int{1, 2, 3, 4, 5})
	fmt.Println(l)
	// res :=
}
