package easy

import "testing"

type treeNode1 struct {
	Val   int
	Left  *treeNode1
	Right *treeNode1
}

// 判断一个二叉树是否是轴对称
func isSymmetric(r *treeNode1) bool {
	if r == nil {
		return true
	}

	//将左右分别遍历写到一个数组对比

	//1 2 4 8 16

	//           0
	//     1              2
	//  3    4         5      6
	// 7 8  9 10     11 12  13  14
	return cmp(r.Left, r.Right)
}

func cmp(p *treeNode1, q *treeNode1) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return cmp(p.Left, q.Right) && cmp(p.Right, q.Left)
}

func TestIsSymmetric(t *testing.T) {
	//用例太麻烦了不想写了
}
