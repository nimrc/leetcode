package _001_two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for x, val := range nums {
		if y, ok := m[target-val]; ok {
			return []int{y, x}
		}

		m[val] = x
	}
	return nil
}
