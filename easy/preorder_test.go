package easy

type treeNode7 struct {
	Val      int
	Children []*treeNode7
}

// N 叉树的前序遍历
func preorder(root *treeNode7) (ans []int) {
	var dfs func(*treeNode7)
	dfs = func(node *treeNode7) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}
