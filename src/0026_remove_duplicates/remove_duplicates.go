package solution

func removeDuplicates(nums []int) int {
	var n = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[n] {
			n++
			nums[n] = nums[i]
		}
	}

	return n + 1
}
