package top100

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := make([]int, 2, 2)
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target && (l == 0 || nums[l-1] < target) {
			ans[0] = l
			break
		}
		if nums[r] == target && nums[r-1] < target {
			ans[0] = r
			break
		}

		mid := l + (r-l)/2
		if mid == l || mid == r {

			ans[0] = -1
			break
		}

		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}

	l, r = 0, len(nums)-1
	for l <= r {
		if nums[r] == target && (r == len(nums)-1 || nums[r+1] > target) {
			ans[1] = r
			break
		}

		if nums[l] == target && nums[l+1] > target {
			ans[1] = l
			break
		}

		mid := l + (r-l)/2
		if mid == l || mid == r {
			ans[1] = -1
			break
		}

		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	return ans
}
