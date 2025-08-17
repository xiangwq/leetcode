package top100

// ListNode 链表公共结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
