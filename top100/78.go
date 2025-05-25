package top100

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	var fn func(data []int, remainNums []int)
	fn = func(data []int, remainNums []int) {
		for k, v := range remainNums {
			var temp = make([]int, len(data))
			copy(temp, data)
			temp = append(temp, v)
			ans = append(ans, temp)
			fn(temp, remainNums[k+1:])
		}
	}
	fn([]int{}, nums)
	return ans
}
