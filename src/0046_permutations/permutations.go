package solution

func permute(nums []int) [][]int {
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

	for i := start; i < end; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		subPermute(nums, start+1, end, res)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
