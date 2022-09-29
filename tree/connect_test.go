package tree

import "testing"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func TestConnect(t *testing.T) {

}

// 解决方法判断
// 前序
// 中序
// 后序
// DFS
// BFS ✅
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		levelCount := len(queue)
		var pre *Node

		for i := 0; i < levelCount; i++ {
			node := queue[0]
			if pre != nil {
				pre.Next = node
			}

			pre = node
			qu := []*Node{}
			if node.Left != nil {
				qu = append(qu, node.Left)
			}
			if node.Right != nil {
				qu = append(qu, node.Right)
			}
			queue = qu
		}
	}

	return root
}
