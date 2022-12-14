package easy

import (
	"fmt"
	"math"
	"testing"
)

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val >= upper || root.Val <= lower {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func TestIsValidBST(t *testing.T) {
	fmt.Println(123)
}
