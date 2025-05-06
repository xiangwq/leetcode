package top100

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld, lm := diameter(root.Left)
	rd, rm := diameter(root.Right)
	return max(ld+rd, max(lm, rm))
}

func diameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lDepth, lMax := diameter(root.Left)
	rDepth, rMax := diameter(root.Right)

	m := max(lDepth+rDepth, max(lMax, rMax))
	return max(lDepth, rDepth) + 1, m
}
