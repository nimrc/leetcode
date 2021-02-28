package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, e := m[nums[i]]; e {
			if i-m[nums[i]] <= k {
				return true
			}
		}
		m[nums[i]] = i
	}

	return false
}
