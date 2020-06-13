package solution

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	res := make([]int, 0)

	for _, v := range nums1 {
		hash[v]++
	}

	for _, v := range nums2 {
		if hash[v] > 0 {
			res = append(res, v)
			hash[v]--
		}

	}

	return res
}
