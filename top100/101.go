package top100

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
func isSymmetric(root *TreeNode) bool {
	return symmetricCompare(root.Left, root.Right)
}

func symmetricCompare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && symmetricCompare(left.Left, right.Right) && symmetricCompare(left.Right, right.Left)
}
