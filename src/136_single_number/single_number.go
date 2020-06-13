package solution

func singleNumber(nums []int) (n int) {
	for i := 0; i < len(nums); i++ {
		n ^= nums[i]
	}
	return
}
