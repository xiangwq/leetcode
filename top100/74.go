package top100

// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)*len(matrix[0])-1
	for l <= r {
		if target == getValue(matrix, l) || target == getValue(matrix, r) {
			return true
		}
		mid := l + (r-l)/2
		if mid == l || mid == r {
			return false
		}
		midValue := getValue(matrix, mid)
		if midValue == target {
			return true
		}
		if midValue < target {
			l = mid
		} else {
			r = mid
		}
	}
	return false
}

func getValue(matrix [][]int, pos int) int {
	x := pos / len(matrix[0])
	y := pos % len(matrix[0])
	return matrix[x][y]
}
