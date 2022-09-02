package easy

import "testing"

type treeNode6 struct {
	Val  int
	Next *treeNode6
}

// 链表的中间结点 利用快慢指针实现
func middleNode(l *treeNode6) *treeNode6 {
	left, right := l, l
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}

	return left
}

func TestMiddleNode(t *testing.T) {
	var l *treeNode6
	middleNode(l)
}
