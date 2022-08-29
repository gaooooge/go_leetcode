package easy

import (
	"fmt"
	"testing"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// 看了大神解法，豁然开朗
func isSameTree(p *treeNode, q *treeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestIssameTree(t *testing.T) {
	var head1 treeNode
	head1.Val = 1

	var p1 treeNode
	p1.Val = 2

	head1.Left = &p1
	var p2 treeNode
	p2.Val = 1

	head1.Right = &p2

	var head2 treeNode
	head2.Val = 1
	var q1 treeNode
	q1.Val = 1
	head2.Left = &q1
	var q2 treeNode
	q2.Val = 2
	head2.Right = &q2

	res := isSameTree(&head1, &head2)

	fmt.Println(res)
}
