package easy


//二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{} //创建返回值的二级数组
    if root == nil { //如果根节点是空则直接返回
        return ret
    }
    q := []*TreeNode{root} //新建树q
    //定义i为0， q长度大于0，i++ //i做为level使用
    for i := 0; len(q) > 0; i++ {
        ret = append(ret, []int{})//加入空数组
        // fmt.Println(ret)
        p := []*TreeNode{} //新建树p
        //定义j为0，j小于q长度，j++
        for j := 0; j < len(q); j++ {
            node := q[j]
            // fmt.Println( node)
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