package easy

import "testing"

type treeNode2 struct {
	Val   int //
	Left  *treeNode2
	Right *treeNode2
}

func maxDepth(r *treeNode2) int {
	if r == nil {
		return 0
	}
	left := maxDepth(r.Left)
	right := maxDepth(r.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func TestMaxDepth(t *testing.T) {

	// maxDepth :=
}
