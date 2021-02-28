package solution

func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
