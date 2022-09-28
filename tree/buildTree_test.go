package tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int       //val
	Left  *TreeNode //left
	Right *TreeNode //right
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rv := postorder[len(postorder)-1]
	var i int
	for i := range inorder {
		if rv == inorder[i] {
			break
		}
	}
	return &TreeNode{
		Val: rv,
		Left: buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	res := buildTree(inorder, postorder)
	fmt.Println(res)
}
