package top100

import (
	"testing"
)

func Test_pathSum(t *testing.T) {
	mockTree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	if sum := pathSum(mockTree, 22); sum != 3 {
		t.Errorf("pathSum() = %v, want %v", sum, 3)
	}
}
