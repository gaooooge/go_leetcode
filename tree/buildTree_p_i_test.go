package tree

import (
	"fmt"
	"testing"
)

func buildTreePI(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	for i := range inorder {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil

}

func TestBuildTreePI(t *testing.T) {
	res := buildTreePI(
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
	)
	fmt.Println(res)
}
