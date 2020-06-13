package solution

import (
	"sort"
)

func containsDuplicateMap(nums []int) bool {
	m := make(map[int]bool, len(nums))
	e := false

	for _, n := range nums {
		if _, e = m[n]; e {
			break
		}
		m[n] = true
	}

	return e
}

func containsDuplicateSort(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

