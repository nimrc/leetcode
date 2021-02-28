package solution

import "sort"

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	var res [][]int

	for n, i := range nums {
		target := 0 - i

		left, right := n+1, len(nums)-1

		if n > 0 && nums[n] == nums[n-1] {
			continue
		}

		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{i, nums[left], nums[right]})

				for left < right && nums[left+1] == nums[left] {
					left++
				}

				for left < right && nums[right-1] == nums[right] {
					right--
				}

				left++
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}

		}
	}

	return res
}
