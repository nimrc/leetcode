package solution

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var solutions = []func(height []int) int{
	// use double pointers
	// T = O(n)
	func(height []int) int {
		max := 0
		start, end := 0, len(height)-1

		for start < end {
			if height[start] < height[end] {
				max = maxInt(max, height[start]*(end-start))
				start++
			} else {
				max = maxInt(max, height[end]*(end-start))
				end--
			}
		}

		return max
	},
}
