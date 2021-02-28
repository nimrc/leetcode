package _011_container_with_most_water

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// use double pointers
// T = O(n)
func maxArea(height []int) int {
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
}
