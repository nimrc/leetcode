package solution

// @link https://leetcode-cn.com/problems/majority-element/

func majorityElement(nums []int) int {
	counter, candidate := 0, 0

	for _, num := range nums {
		if counter == 0 {
			candidate = num
		}

		if candidate == num {
			counter++
		} else {
			counter--
		}
	}

	return candidate
}
