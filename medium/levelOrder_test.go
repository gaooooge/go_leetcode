package medium

type treeNode8 struct {
	Val   int
	Left  *treeNode8
	Right *treeNode8
}

// 二叉树的层序遍历
func levelOrder(root *treeNode8) [][]int {
	ret := [][]int{}
	if root != nil {
		return ret
	}

	q := []*treeNode8{root}

	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*treeNode8{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	return ret
}
