package medium

import (
	"fmt"
	"testing"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// 最长同值路径
func longestUnivaluePath(root *treeNode) (ans int) {
	var dfs func(*treeNode) int
	dfs = func(node *treeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right1 = right + 1
		}
		ans = max(ans, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func TestLongestUnivaluePath(t *testing.T) {
	var l *treeNode
	res := longestUnivaluePath(l)
	fmt.Println(res)
}
