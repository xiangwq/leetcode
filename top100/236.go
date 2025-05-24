package top100

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := findLowestCommonAncestor(root, p, q)
	return node
}

func findLowestCommonAncestor(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	l, ok1 := findLowestCommonAncestor(root.Left, p, q)
	if l != nil {
		return l, true
	}
	r, ok2 := findLowestCommonAncestor(root.Right, p, q)
	if r != nil {
		return r, true
	}

	if (root.Val == p.Val || root.Val == q.Val) && (ok1 || ok2) {
		return root, true
	}
	if ok1 && ok2 {
		return root, true
	}

	if (root.Val == p.Val || root.Val == q.Val) || (ok1 || ok2) {
		return nil, true
	}
	return nil, false
}
