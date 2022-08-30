package easy

type treeNode4 struct {
	Val   int //
	Left  *treeNode4
	Right *treeNode4
}

// 判断平衡树
func isBalanced(r *treeNode4) bool {
	if r == nil {
		return true
	}

	return abs(height(r.Left)-height(r.Right)) <= 1 && isBalanced(r.Left) && isBalanced(r.Right)
}

func height(r *treeNode4) int {
	if r == nil {
		return 0
	}
	return max(height(r.Left), height(r.Right)) + 1

}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {

	if x < 0 {
		return -1 * x
	}

	return x
}
