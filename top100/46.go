package top100

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	answer := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		remain := make([]int, 0)
		remain = append(remain, nums[0:i]...)
		if i < len(nums)-1 {
			remain = append(remain, nums[i+1:]...)
		}
		tempAnswer := permute(remain)
		for _, v := range tempAnswer {
			temp := []int{nums[i]}
			temp = append(temp, v...)
			answer = append(answer, temp)
		}
	}
	return answer
}
