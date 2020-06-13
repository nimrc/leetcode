package solution

func permuteUnique(nums []int) [][]int {
	var res [][]int
	subPermute(nums, 0, len(nums), &res)

	return res
}

func subPermute(nums []int, start, end int, res *[][]int) {
	if start == end {
		item := make([]int, end)
		copy(item, nums)

		*res = append(*res, item)
		return
	}

	used := make(map[int]bool)

	for i := start; i < end; i++ {
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		nums[i], nums[start] = nums[start], nums[i]

		subPermute(nums, start+1, end, res)

		nums[i], nums[start] = nums[start], nums[i]
	}
}
