package top100

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
func pathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	return dfs(root, targetSum, prefixSum, 0)
}

func dfs(root *TreeNode, targetSum int, pm map[int]int, sum int) int {
	if root == nil {
		return 0
	}
	sum = root.Val + sum
	pm[sum]++
	ln := dfs(root.Left, targetSum, pm, sum)
	rn := dfs(root.Right, targetSum, pm, sum)
	pm[sum]--
	num := ln + rn + pm[sum-targetSum]
	return num
}
