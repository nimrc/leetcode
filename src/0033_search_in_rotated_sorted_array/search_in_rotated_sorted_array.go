package solution

// @link https://leetcode-cn.com/problems/search-in-rotated-sorted-array/submissions/

func search(nums []int, target int) int {
	n := len(nums) - 1
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if target < nums[mid] && target >= nums[0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[n] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
