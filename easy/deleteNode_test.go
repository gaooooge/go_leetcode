package easy

//删除链表中的节点  删除节点只需要取代下一个节点即可
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}